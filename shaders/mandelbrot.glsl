
#version 410 core

#include <uniforms.glsl>

#include <complex.glsl>
#include <fractals.glsl>

out vec4 color;

const float fractalMaxX = 1.;

const float fractalMinX = -1.;
const float fractalMaxY = 1.0;
const float fractalMinY = -1.0;
float scaleX = fractalMaxX - fractalMinX;
float scaleY = fractalMaxY - fractalMinY;

vec3 hsv2rgb(vec3 c)
{
   float H = (1. - abs(mod((c.x / 60.0), 2.0) - 1.));
   float S = c.y;
   float V = c.z;


   float C = V * S; // V * S
   float X = C * H; // C * H
   float m = V - C;

   vec3 rgb;

   if (H >= 0. && H < 60.) {
      rgb = vec3(C, X, 0);
   } else if (H >= 60. && H < 120.) {
      rgb = vec3(X, C, 0);
   } else if (H >= 120. && H < 180.) {
      rgb = vec3(0, C, X);
   } else if (H >= 180. && H < 240.) {
      rgb = vec3(0, X, C);
   } else if (H >= 240. && H < 300.) {
      rgb = vec3(X, 0, C);
   } else if (H >= 300. && H < 360.) {
      rgb = vec3(C, 0, X);
   }

   rgb += m;
   return rgb;
}


void main() { 
   // Normalize the current screen coordinates between the range -2.5 and 1 for x
   // and -1 and 1 
   float scaledX = (gl_FragCoord.x/Resolution.x) * scaleX + fractalMinX;
   float scaledY = (gl_FragCoord.y/Resolution.y) * scaleY + fractalMinY;

   float z = Zoom;
   if (z < 0) {
      z = 1;
   }

   vec2 uv = vec2((scaledX / z) , -(scaledY/ z) );

   uv.x += Pan.x;
   uv.y += Pan.y;

//vec2(−0.8, 0.156)
   MandelbrotResult result = mandelbrot(MaxIterations, uv,  vec2(-0.,0.));
   
   int iterations = result.iterations;


   float H = mod(360. * log2(float(iterations)), 360.);
   float S = 1.;
   float V =  sqrt(float(iterations) / 255.); 
   float col = cos(float(iterations) / 255. );

   vec3 rgb = hsv2rgb(vec3(H, S, V));
   // if (iterations < min(Time, 3000)) {
      color = vec4(rgb, 1.); 
   // } else {
   //    color = vec4(0.);
   // }

   // color = vec4(1., .5, .25, 1.);
}