
#version 410 core

precision highp double;

#include <uniforms.glsl>

#include <complex.glsl>
#include <fractals.glsl>

//int[] pallette = int[] (0xF94144, 0xF3722C, 0xF8961E, 0xF9844A, 0xF9C74F, 0x90BE6D, 0x43AA8B, 0x4D908E, 0x577590, 0x277DA1);

int[] pallette = int[](0x757BC8, 0x8E94F2, 0xBBADFF, 0xDAB6FC, 0xE0C3FC);

uniform int[256] palette; 

vec3 rgb(int hex) {
   float red = float((hex >> 16) & 0xFF) / 255.;
   float green = float((hex >> 8) & 0xFF) / 255.;
   float blue = float((hex) & 0xFF) / 255.;

   return vec3(red, green, blue);
}

vec3 colorAt(float index) {
   int mul = int(index);
   vec3 color1 = rgb(pallette[mul  ]);
   vec3 color2 = rgb(pallette[mul + 1]);

   vec3 pct = vec3(index);

   // pct.r = smoothstep(0.0, 1.0, index);
   // pct.g = sin(index * 3.14159265359 );
   // pct.b = pow(index, 0.5);
   return mix(color1, color2, pct);
}

out vec4 color;

const float fractalMaxX = 1.;

const float fractalMinX = -1.;
const float fractalMaxY = 1.0;
const float fractalMinY = -1.0;
float scaleX = fractalMaxX - fractalMinX;
float scaleY = fractalMaxY - fractalMinY;

vec3 hsv2rgb(vec3 c) {
   float H = (1. - abs(mod((c.x / 60.0), 2.0) - 1.));
   float S = c.y;
   float V = c.z;

   float C = V * S; // V * S
   float X = C * H; // C * H
   float m = V - C;

   vec3 rgb;

   if(H >= 0. && H < 60.) {
      rgb = vec3(C, X, 0);
   } else if(H >= 60. && H < 120.) {
      rgb = vec3(X, C, 0);
   } else if(H >= 120. && H < 180.) {
      rgb = vec3(0, C, X);
   } else if(H >= 180. && H < 240.) {
      rgb = vec3(0, X, C);
   } else if(H >= 240. && H < 300.) {
      rgb = vec3(X, 0, C);
   } else if(H >= 300. && H < 360.) {
      rgb = vec3(C, 0, X);
   }

   rgb += m;
   return rgb;
}

void main() { 
   // Normalize the current screen coordinates between the range -2.5 and 1 for x
   // and -1 and 1 
   float scaledX = (gl_FragCoord.x / Resolution.x) * scaleX + fractalMinX;
   float scaledY = (gl_FragCoord.y / Resolution.y) * scaleY + fractalMinY;

   float z = Zoom;
   if(z < 0) {
      z = 1;
   }

   vec2 uv = vec2((scaledX / z), -(scaledY / z));

   uv.x += Pan.x;
   uv.y += Pan.y;

//vec2(−0.8, 0.156)
   MandelbrotResult result = mandelbrot(MaxIterations, uv, vec2(-0., 0.));

   int iterations = result.iterations;

   // float H = mod(360. * log2(float(iterations)), 360.);
   // float S = 1.;
   // float V =  sqrt(float(iterations) / 255.); 
   // float col = cos(float(iterations) / 255. );

   vec3 pct = rgb(palette[ int( float(iterations)/float(MaxIterations) * 255.      )  ]);

   if(iterations == MaxIterations) {
      color = vec4(0.);
   } else {
      color = vec4(pct, 0.);
   }
}