using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Robot
{
    internal class Robot
    {
        private string topic = "topic/";
        private string controlTopic;
        private string id;
        private State state;
        private Broker broker;
        private string error;
       
        public Robot(string id, Broker broker) 
        {
            state = State.IDLE;
            this.id = "ROBOT_" + id;
            topic += this.id;
            controlTopic = "topic/" + this.id + "/signal";
            this.broker = broker;
            error = "code 1337 message ¯\\_(ツ)_/¯";
            
        }

        public void Run()
        {
            broker.Connect().Wait();
            broker.Message("topic/robot/new", id);
            broker.Subscribe(controlTopic, m => 
            {
                Console.WriteLine(m);
                switch(m)
                {
                    case "START":
                        state = State.RUNNING;
                        OnStart();
                        break;
                    case "HALT":
                        state = State.IDLE;
                        OnStop();
                        break;
                    default:
                        state = State.FAILED;
                        OnFailure();
                        break;
                }
            });
            int step = 0;
            int maxSteps = 5;
            GPU.Builder builder = new();
            while (true) 
            {
                Thread.Sleep(1000);
                if (state == State.RUNNING) 
                {
                    switch (step)
                    {
                        case 0:
                            builder.Pcb(GPU.Parts.NewPcb());
                            broker.Message("topic/production/gpu", "pcb");
                            break;
                        case 1:
                            builder.Processor(GPU.Parts.NewProcessor());
                            broker.Message("topic/production/gpu", "processor");
                            break;
                        case 2:
                            builder.Fan(GPU.Parts.NewFan());
                            broker.Message("topic/production/gpu", "fan");
                            break;
                        case 3:
                            builder.Firmware(GPU.Parts.firmware);
                            broker.Message("topic/production/gpu", "firmware");
                            break;
                        case 4:
                            GPU.GPU gpu = builder.Build();
                            broker.Message("topic/production/gpu/completed", $"robot={id},{gpu.Package()}");
                            builder = new();
                            break;
                    }
                    broker.Message(topic, state.ToString());
                    step = (step + 1) % maxSteps;
                }
                
            }
        }

        private void OnStart()
        {
            broker.Message(topic, "STARTING");
        }

        private void OnStop()
        {
            
            broker.Message(topic, "STOPPING");
        }

        private void OnFailure()
        {
            broker.Message(topic, error);

        }
    }


    internal enum State
    {
        IDLE,
        STARTING,
        RUNNING,
        STOPPING,
        FAILED
    }
}
