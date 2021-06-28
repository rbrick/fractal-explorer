#define BURNING_SHIP Mode == 1

vec2 square(vec2 v, bool burning_ship) {
   if (burning_ship) {
      v = vec2(abs(v.x), abs(v.y));   
   }
   return cmplx_multiply(v, v);
}

int mandelbrot(int iter, vec2 pos) {
   vec2 z = vec2(0, 0);
   vec2 c = vec2(pos.x, pos.y);
   
   for (int i = 0; i < iter; i++) {
     vec2 v = vec2(z.x, z.y);
     z = square(v, BURNING_SHIP);
     z = cmplx_add(z,c);
     if (z.x*z.x +z.y*z.y >= 2.*2.){
        return i;
     }
   }
   return iter;
}


int julia() {
   return 0;
}