using System.Text;
using MQTTnet;
using MQTTnet.Client;

namespace Robot;

internal class Broker
{
    public delegate void OnMessage(string message);

    private readonly IMqttClient client;
    private readonly MqttFactory factory;

    public Broker()
    {
        factory = new MqttFactory();
        client = factory.CreateMqttClient();
    }

    public async Task Connect()
    {
        var options = new MqttClientOptionsBuilder().WithTcpServer("127.0.0.1", 1883).Build();
        Console.WriteLine("Connecting to broker ...");
        await client.ConnectAsync(options, CancellationToken.None);
        Console.WriteLine("Connection established");
    }

    public async void Message(string topic, string message)
    {
        var mqttMessage = new MqttApplicationMessageBuilder().WithTopic(topic).WithPayload(message).Build();
        await client.PublishAsync(mqttMessage);
    }

    public async void Subscribe(string topic, OnMessage onMessage)
    {
        var mqttTopic = factory.CreateSubscribeOptionsBuilder().WithTopicFilter(t => t.WithTopic(topic)).Build();
        Console.WriteLine("Subscribing to topic [" + topic + "] ...");
        await client.SubscribeAsync(mqttTopic, CancellationToken.None);
        client.ApplicationMessageReceivedAsync += e =>
        {
            var payload = Encoding.Default.GetString(e.ApplicationMessage.PayloadSegment);
            onMessage(payload);
            return Task.CompletedTask;
        };
        Console.WriteLine("Subscribed");
    }
}
