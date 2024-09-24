using System;
using System.Drawing;
using System.Drawing.Imaging;

public class ACaptcha
{
    private static Random _random = new Random();

    public static string GenerateCaptcha(int width, int height)
    {
        // Generate a random code
        string code = GenerateRandomCode(5);

        // Create a bitmap image
        Bitmap image = new Bitmap(width, height);
        Graphics graphics = Graphics.FromImage(image);

        // Draw the code on the image
        graphics.DrawString(code, new Font("Arial", 24), Brushes.Black, 10, 10);

        // Add some noise to the image
        for (int i = 0; i < 10; i++)
        {
            graphics.DrawLine(Pens.Black, _random.Next(width), _random.Next(height), _random.Next(width), _random.Next(height));
        }

        // Save the image to a byte array
        byte[] imageData = new byte[0];
        using (MemoryStream stream = new MemoryStream())
        {
            image.Save(stream, ImageFormat.Png);
            imageData = stream.ToArray();
        }

        return Convert.ToBase64String(imageData);
    }

    private static string GenerateRandomCode(int length)
    {
        string characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
        char[] code = new char[length];
        for (int i = 0; i < length; i++)
        {
            code[i] = characters[_random.Next(characters.Length)];
        }
        return new string(code);
    }
}
