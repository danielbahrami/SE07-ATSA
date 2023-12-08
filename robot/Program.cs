internal class Program
{
    private static void Main(string[] args)
    {
        string id = "DEFAULT";
        if (args.Length > 0) { id = args[0]; }
        Robot.Broker b = new();
        Robot.Robot r = new(id, b);
        r.Run();
    }
}