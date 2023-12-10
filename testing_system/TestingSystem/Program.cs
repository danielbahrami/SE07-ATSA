internal class Program
{
    private static void Main(string[] args)
    {
        bool run = true;
        TestingSystem.Broker broker = new TestingSystem.Broker();
        TestingSystem.Tester tester = new TestingSystem.Tester();
        broker.Connect().Wait();

        broker.Subscribe("topic/production/gpu/completed", m =>
        {
            //True/False for test results
            broker.Message("topic/testing/gpu/completed", tester.TestProduct(m).ToString());
        });

        while (run)
        {

        }
    }
}