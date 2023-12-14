using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace GPU
{
    internal class GPU
    {
        private string pcb;
        private string processor;
        private string fan;
        private string firmware;

        internal GPU()
        {
            
        }

        internal void SetPcb(string pcb) { this.pcb = pcb; }
        internal void SetProcessor(string processor) { this.processor = processor; }
        internal void SetFan(string fan) {  this.fan = fan; }
        internal void SetFirmware(string firmware) {  this.firmware = firmware; }

        public string Package()
        {
            return $"pcb={pcb},processor={processor},fan={fan},firmware={firmware},integrity={Parts.Integrity()}";
        }
    }

    internal class Parts
    {
        private static int pcbs = 0;
        public static string NewPcb()
        {
            return $"pcb_{pcbs++}";
        }

        private static int processors = 0;
        public static string NewProcessor()
        {
            return $"processor_{processors++}";
        }

        private static int fans = 0;
        public static string NewFan() 
        {
            return $"fan_{fans++}";
        }

        public static readonly string firmware = "v0.0.1";

        // Generate a random number that will determin if the product will pass testing
        private static readonly Random radnom = new Random();
        internal static int Integrity()
        {
            return radnom.Next(0,100);
        }
    }

    internal class Builder
    {
        private GPU gpu;
        public Builder() 
        {
            this.gpu = new();
        }

        public Builder Pcb(string pcb) 
        {
            gpu.SetPcb(pcb); return this;
        }
        public Builder Processor(string processor)
        {
            gpu.SetProcessor(processor); return this;
        }
        public Builder Fan(string fan)
        {
            gpu.SetFan(fan); return this;
        }
        public Builder Firmware(string firmware)
        {
            gpu.SetFirmware(firmware); return this;
        }

        public GPU Build() 
        {
            return gpu;
        }

    }
}
