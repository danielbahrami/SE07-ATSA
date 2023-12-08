using Robot;

internal class Program
{
    private static void Main(string[] args)
    {
        var id = "DEFAULT";
        if (args.Length > 0) id = args[0];
        Broker b = new();
        Robot.Robot r = new(id, b);
        r.Run();
    }
}
