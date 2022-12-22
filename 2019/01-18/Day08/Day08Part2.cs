using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day08Part2
    {
        class Image
        {
            string data;
            int width;
            int height;
            int depth;

            public Image(string data, int width, int height)
            {
                this.data = data;
                this.width = width;
                this.height = height;
                this.depth = data.Length / (width * height);
            }

            public int getPixel(int x, int y, int z)
            {
                return int.Parse(data[z * width * height + y * width + x] + "");
            }

            public int getPixel(int x, int y)
            {
                for (var z=0; z<depth; z++)
                {
                    var color = getPixel(x, y, z);
                    if (color != 2)
                        return color;
                }

                return 0;
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("08");
            var width = 25;
            var height = 6;

            var image = new Image(input, width, height);

            for (var y = 0; y < height; y++)
            {
                for (var x = 0; x < width; x++)
                    Console.Write(image.getPixel(x, y) == 1 ? '#' : ' ');

                Console.WriteLine();
            }
        }
    }
}
