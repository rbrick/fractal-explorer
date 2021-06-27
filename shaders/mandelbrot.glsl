
#version 150 core

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

void main() { 
   // Normalize the current screen coordinates between the range -2.5 and 1 for x
   // and -1 and 1 
   float scaledX = (gl_FragCoord.x/Resolution.x) * scaleX + fractalMinX;
   float scaledY = (gl_FragCoord.y/Resolution.y) * scaleY + fractalMinY;

   float z = Zoom;
   if (z < 0) {
      z = 1;
   }

   vec2 uv = vec2((scaledX / z) , (scaledY/ z) );

   uv.x += Pan.x;
   uv.y += Pan.y;

   int iterations = mandelbrot(MaxIterations, uv);

   if (iterations < MaxIterations) {
      color = vec4(vec3(float(iterations)/float(MaxIterations)), 1.); 
   } else {
      color = vec4(1.);
   }

   // color = vec4(1., .5, .25, 1.);
}