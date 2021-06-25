
#version 150 core

#include <complex.glsl>

const float fractalMaxX = 1.5;
const float fractalMinX = -1.5;
const float fractalMaxY = 1.0;
const float fractalMinY = -1.0;
float scaleX = fractalMaxX - fractalMinX;
float scaleY = fractalMaxY - fractalMinY;

int julia(int iter, vec2 pos) {
   vec2 z = vec2(0.,0.);
   vec2 c = vec2(pos.x, pos.y);
   
   for (int i = 0; i < iter; i++) {
     vec2 absol = vec2(abs(z.x), abs(z.y));
     z = cmplx_multiply(absol, absol);
     z = cmplx_add(z,c);
     if (z.x*z.x +z.y*z.y >= 2.*2.){
        return i;
     }
   }
   return iter;
}
void mainImage( out vec4 fragColor, in vec2 fragCoord )
{

   //  float zoom = 12.2 + float(iFrame)/(iTime*4.);
   //  // Normalized pixel coordinates (from 0 to 1)
   //  float scaledX = (fragCoord.x/iResolution.x) * scaleX + fractalMinX;
   //  float scaledY = (fragCoord.y/iResolution.y) * scaleY + fractalMinY;
   //  vec2 uv = vec2(scaledX / zoom - (1.75), -(scaledY / zoom) - (0.02));
   //  // Time varying pixel color
   //  vec3 col = vec3(0.,0.,0.);
   //  int iterations = julia(200, uv);
   //  if (iterations < 200) {
   //     float percentage = float(iterations)/360.;
       
       
   //     // it escaped
   //     col = vec3(percentage);
   //  } else {
   //     col = vec3(1.,1., 1.);
   //  }
   //  // Output to screen
   //  fragColor = vec4(col,1.0);
    
   //  zoom++;
}