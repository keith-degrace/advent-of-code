using System;
using System.Collections.Generic;
using System.IO;
using System.Reflection;

namespace Advent_of_Code_2019.Utils
{
    class InputLoader
    {
        public static string loadAsString(String day)
        {
            return System.IO.File.ReadAllText(getFileName(day));
        }
        public static string[] loadAsStringArray(String day)
        {
            return System.IO.File.ReadAllLines(getFileName(day));
        }
        private static string getFileName(String day)
        {
            return "../../../Day" + day + "/Day" + day + ".txt";
        }
    }
}
