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
            while (true) 
            {
                Thread.Sleep(1000);
                broker.Message(topic, state.ToString());
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
