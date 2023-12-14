using System;
using System.IO;

internal class Program
{
    private static void Main(string[] args)
    {
        bool run = true;
        string brokerAddress = Environment.GetEnvironmentVariable("BROKER");

        TestingSystem.Broker broker = new(brokerAddress);
        TestingSystem.Tester tester = new();
        broker.Connect().Wait();

        broker.Subscribe("topic/production/gpu/completed", m =>
        {
            broker.Message("topic/testing/gpu/completed", tester.TestProduct(m).ToString());
        });

        while (run)
        {

        }
    }
}
