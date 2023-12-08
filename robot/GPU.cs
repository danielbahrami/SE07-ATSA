namespace GPU;

internal class GPU
{
    private string fan;
    private string firmware;
    private string pcb;
    private string processor;

    internal void SetPcb(string pcb)
    {
        this.pcb = pcb;
    }

    internal void SetProcessor(string processor)
    {
        this.processor = processor;
    }

    internal void SetFan(string fan)
    {
        this.fan = fan;
    }

    internal void SetFirmware(string firmware)
    {
        this.firmware = firmware;
    }

    public string Package()
    {
        return $"pcb={pcb},processor={processor},fan={fan},firmware={firmware}";
    }
}

internal class Parts
{
    private static int pcbs;

    private static int processors;

    private static int fans;

    public static string firmware = "v0.0.1";

    public static string NewPcb()
    {
        return $"pcb_{pcbs++}";
    }

    public static string NewProcessor()
    {
        return $"processor_{processors++}";
    }

    public static string NewFan()
    {
        return $"fan_{fans++}";
    }
}

internal class Builder
{
    private readonly GPU gpu;

    public Builder()
    {
        gpu = new GPU();
    }

    public Builder Pcb(string pcb)
    {
        gpu.SetPcb(pcb);
        return this;
    }

    public Builder Processor(string processor)
    {
        gpu.SetProcessor(processor);
        return this;
    }

    public Builder Fan(string fan)
    {
        gpu.SetFan(fan);
        return this;
    }

    public Builder Firmware(string firmware)
    {
        gpu.SetFirmware(firmware);
        return this;
    }

    public GPU Build()
    {
        return gpu;
    }
}
