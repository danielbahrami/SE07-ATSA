using MQTTnet;
using MQTTnet.Client;
using MQTTnet.Server;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace NotificationSystem
{
    internal class Broker
    {
        private IMqttClient client;
        private MqttFactory factory;

        public Broker()
        {
            factory = new MqttFactory();
            client = factory.CreateMqttClient();
        }

        public async Task Connect() 
        {
            var options = new MqttClientOptionsBuilder().WithTcpServer("127.0.0.1",1883).Build();
            Console.WriteLine("Connecting to broker ...");
            await client.ConnectAsync(options, CancellationToken.None);
            Console.WriteLine("Connection established");
        }

        public async void Message(string topic, string message)
        {
            var mqttMessage = new MqttApplicationMessageBuilder().WithTopic(topic).WithPayload(message).Build();
            await client.PublishAsync(mqttMessage);
        }

        public delegate void OnMessage(string message);

        public async void Subscribe(string topic, OnMessage onMessage) 
        {
            var mqttTopic = factory.CreateSubscribeOptionsBuilder().WithTopicFilter(t => t.WithTopic(topic)).Build();
            Console.WriteLine("Subscribing to topic [" + topic +"] ...");
            await client.SubscribeAsync(mqttTopic, CancellationToken.None);
            client.ApplicationMessageReceivedAsync += e =>
            {
                string payload = Encoding.Default.GetString(e.ApplicationMessage.PayloadSegment);
                onMessage(payload);
                return Task.CompletedTask;
            };
            Console.WriteLine("Subscribed");
        }
    }
}
