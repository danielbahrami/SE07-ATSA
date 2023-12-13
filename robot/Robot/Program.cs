using System;
using System.IO;

internal class Program
{
    private static void Main(string[] args)
    {
        string id = Environment.GetEnvironmentVariable("ID");
        Console.WriteLine("Starting robot {0} ...", id);

        string brokerAddress = Environment.GetEnvironmentVariable("BROKER");
        Robot.Broker b = new(brokerAddress);
        Robot.Robot r = new(id, b);
        r.Run();
    }
}