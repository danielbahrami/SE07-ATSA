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
        private string notifyTopic = "topic/robot/notify";
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
            error = "FAILED: code 1337 message ¯\\_(ツ)_/¯";
            
        }

        public void Run()
        {
            broker.Connect().Wait();
            broker.Message("topic/robot/new", id);
            broker.Subscribe(controlTopic, m => 
            {
                Console.WriteLine(m);
                switch(m.ToUpper())
                {
                    case "START":
                        OnStart();
                        break;
                    case "HALT":
                        OnStop();
                        break;
                    default:
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
                    step = (step + 1) % maxSteps;
                }
                broker.Message(topic, state.ToString());
            }
        }

        private void OnStart()
        {
            Console.WriteLine("Starting ...");
            this.state = State.STARTING;
            broker.Message(notifyTopic, $"robot_id={this.id},state={this.state}");
            Thread.Sleep(2000);
            this.state = State.RUNNING;
            Console.WriteLine("Running ...");
            broker.Message(notifyTopic, $"robot_id={this.id},state={this.state}");
        }

        private void OnStop()
        {
            Console.WriteLine("Stopping ...");
            this.state = State.STOPPING;
            broker.Message(notifyTopic, $"robot_id={this.id},state={this.state}");
            Thread.Sleep(2000);
            this.state = State.IDLE;
            Console.WriteLine("Idle ...");
            broker.Message(notifyTopic, $"robot_id={this.id},state={this.state}");
        }

        private void OnFailure()
        {
            this.state = State.FAILED;
            Console.WriteLine("Failed ...");
            broker.Message(notifyTopic, $"robot_id={this.id},state={this.state},error={error}");
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
