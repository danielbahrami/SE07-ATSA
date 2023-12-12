using GPU;

namespace Robot;

internal class Robot
{
    private readonly Broker broker;
    private readonly string controlTopic;
    private readonly string error;
    private readonly string id;
    private readonly string topic = "topic/";
    private State state;

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
            switch (m)
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
        var step = 0;
        var maxSteps = 5;
        Builder builder = new();
        while (true)
        {
            Thread.Sleep(1000);
            if (state == State.RUNNING)
            {
                switch (step)
                {
                    case 0:
                        builder.Pcb(Parts.NewPcb());
                        broker.Message("topic/production/gpu", "pcb");
                        break;
                    case 1:
                        builder.Processor(Parts.NewProcessor());
                        broker.Message("topic/production/gpu", "processor");
                        break;
                    case 2:
                        builder.Fan(Parts.NewFan());
                        broker.Message("topic/production/gpu", "fan");
                        break;
                    case 3:
                        builder.Firmware(Parts.firmware);
                        broker.Message("topic/production/gpu", "firmware");
                        break;
                    case 4:
                        var gpu = builder.Build();
                        broker.Message("topic/production/gpu/completed", $"robot={id},{gpu.Package()}");
                        builder = new Builder();
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
