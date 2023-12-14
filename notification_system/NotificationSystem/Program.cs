using System;
using System.IO;
using NotificationSystem;
using static System.Net.Mime.MediaTypeNames;

internal class Program
{
    // Event Queue
    private static Queue<string> queue = new Queue<string>();

    private static void Main(string[] args)
    {
        string brokerAddress = Environment.GetEnvironmentVariable("BROKER");

        Broker broker = new(brokerAddress);
        broker.Connect().Wait();
        broker.Subscribe("topic/robot/notify", m =>
        {
            Console.WriteLine("Notifying ... {0}", m);
            queue.Enqueue(m);
        });

        var builder = WebApplication.CreateBuilder(args);
        builder.Services.AddCors().AddHttpContextAccessor();
                
        var app = builder.Build();
        app.UseCors(p => p.AllowAnyHeader().AllowAnyMethod().WithOrigins("http://localhost:3000"));

        // Subscribe for events
        app.MapGet("/sse", async (IHttpContextAccessor accessor, CancellationToken cancellationToken) =>
        {
            var res = accessor.HttpContext!.Response;
            Console.WriteLine("Registering ... ");

            // Header for SSE
            res.Headers["Content-Type"] = "text/event-stream";
          
            string payload = "";
            // Keep running
            while (!cancellationToken.IsCancellationRequested)
            {
                await foreach (var e in EventQueue())
                {
                    // Send event to subscribers
                    Console.WriteLine("Sending ...");
                    await res.WriteAsync($"data: {e}\n\n", cancellationToken); // Must be structured like: "data: <text>\n\n"
                    await res.Body.FlushAsync(cancellationToken);
                }
            }
            
        });

        // Publish events (for testing)
        app.MapGet("/sse/invoke", (IHttpContextAccessor accessor) =>
        {
            var req = accessor.HttpContext!.Response;
            Console.WriteLine("Notifying ...");

            var payload = "{\"message\":\"new\"}";
            queue.Enqueue(payload);
            req.WriteAsJsonAsync(payload);
        });
        
        app.Run();
    }

    // Async queue
    public static async IAsyncEnumerable<string> EventQueue()
    {
        if (queue.Count > 0)
        {
            yield return queue.Dequeue();
        }
    }
}
