#define BURNING_SHIP Mode == 1


vec2 square(vec2 v, bool burning_ship) {
   if (burning_ship) {
      v = vec2(abs(v.x), abs(v.y));   
   }
   return cmplx_multiply(v, v);
}

struct MandelbrotResult {
   int iterations;
   float dist;
};

float mag(vec2 v) {
   return abs(sqrt( v.x*v.x  + v.y*v.y));
}

MandelbrotResult mandelbrot(int iter, vec2 pos, vec2 center) {
   vec2 z = vec2(0., 0.);
   vec2 c = vec2(pos.x, pos.y);

   vec2 dz = vec2(0., 0.); // define a distance 

   int result = iter;
   
   for (int i = 0; i < iter; i++) {
     vec2 v = vec2(z.x,z.y);
     z = square(v, BURNING_SHIP);
     z = cmplx_add(z,c); // z^2 + c
   
   //   dz = (2 * z) * (dz + 1);
      vec2 temp = cmplx_multiply(vec2(2,0), z);
      dz = cmplx_multiply(temp, dz);
      dz = cmplx_add(dz, vec2(1,0));
 

     if (z.x*z.x +z.y*z.y >=2.0*2.0){
        result = i; 
        break;
     }
   }


   // some result
   
   float mz = mag(z);
   float mdz = mag(dz);
   float dist = log(mz*mdz) * mz / mdz; 
   
   return MandelbrotResult(result, dist);
}


int julia() {
   return 0;
}