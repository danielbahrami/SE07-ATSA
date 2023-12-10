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

        // Some product test :-|
        // ProductString should look like this `robot={id},pcb={pcb},processor={processor},fan={fan},firmware={firmware},integrity={integrity}`
        public bool TestProduct(string productString)
        {
            Console.Write("Testing {0} ... ", productString);
            string[] lines = productString.Split(',');
            if (lines.Length == 0) 
            {
                Console.WriteLine("NO LINES!!!");
                return false;
            }

            bool pass = true;

            foreach (string line in lines)
            {
                // Check for currect firmware
                if (line.StartsWith("firmware"))
                {
                    string[] lineSplit = line.Split("=");
                    Console.Write("{0} ... ", lineSplit[1] == firmwareVersion);
                    pass &= lineSplit[1] == firmwareVersion;
                }

                // Check for integrity
                else if (line.StartsWith("integrity"))
                {
                    string[] lineSplit = line.Split("=");
                    int integrity = Int32.Parse(lineSplit[1]);
                    Console.Write("{0} ... ", (integrity > 15 && integrity < 85));
                    pass &= (integrity > 15 && integrity < 85); // idk some shit ...
                }
            }
            Console.WriteLine(pass);
            return pass;
        }
    }


}
