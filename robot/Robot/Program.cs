using System;
using System.IO;

internal class Program
{
    private static void Main(string[] args)
    {
        Console.WriteLine("Starting robot ...");
        string id = "DEFAULT";
        if (args.Length > 0) { id = args[0]; }
        string brokerAddress = Environment.GetEnvironmentVariable("BROKER");
        Robot.Broker b = new(brokerAddress);
        Robot.Robot r = new(id, b);
        r.Run();
    }
}