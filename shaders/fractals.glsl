int mandelbrot(int iter, vec2 pos) {
   vec2 z = vec2(pos.x,pos.y);
   vec2 c = vec2(pos.x, pos.y);
   
   for (int i = 0; i < iter; i++) {
     vec2 absol = vec2(z.x, z.y);
     z = cmplx_multiply(absol, absol);
     z = cmplx_add(z,c);
     if (z.x*z.x +z.y*z.y >= 2.*2.){
        return i;
     }
   }
   return iter;
}