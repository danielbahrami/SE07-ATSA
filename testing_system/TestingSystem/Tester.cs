using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TestingSystem
{
    internal class Tester
    {
        public Tester()
        {

        }

        private static readonly string firmwareVersion = "0.0.1";
        
        public string TestProduct(string productString)
        {
            Console.Write("Testing {0} ... ", productString);
            string[] lines = productString.Split(',');
            if (lines.Length == 0)
            {
                Console.WriteLine("NO LINES!!!");
                return "final:False";
            }

            bool pass = true;
            string report = "";

            foreach (string line in lines)
            {
                // Check for currect firmware
                if (line.StartsWith("firmware"))
                {
                    string[] lineSplit = line.Split("=");
                    Console.Write("{0} ... ", lineSplit[1] == firmwareVersion);
                    report += "firmware:" + (lineSplit[1] == firmwareVersion).ToString();
                    pass &= lineSplit[1] == firmwareVersion;
                }

                // Check for integrity
                else if (line.StartsWith("integrity"))
                {
                    string[] lineSplit = line.Split("=");
                    int integrity = Int32.Parse(lineSplit[1]);
                    Console.Write("{0} ... ", (integrity > 15 && integrity < 85));
                    report += "integrity:" + (integrity > 15 && integrity < 85).ToString();
                    pass &= (integrity > 15 && integrity < 85);
                }
            }
            report += "final:" + pass.ToString();
            Console.WriteLine(report);
            return report;
        }
    }
}
