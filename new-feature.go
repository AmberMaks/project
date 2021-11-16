using System;

namespace practical3__1
{
    class Program
    {
        static void Main(string[] args)
        {
            var massiv = new int[3, 3];
            var elem = new int[3];
            int k = 0;
            for (int i = 0; i < 3; i++)
            {
                for (int j = 0; j < 3; j++)
                {
                    massiv[i, j] = int.Parse(Console.ReadLine());
                    if (i == j)
                    {
                        elem[k] = massiv[i, j];
                        k += 1;
                    }
                }
            }
            for (int z = 0; z < 2; z++)
            {
                Console.WriteLine($"{elem[z]}");
            }
        }
    }
}
