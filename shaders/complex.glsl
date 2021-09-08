// This is a simple library for working with complex numbers

//performs a complex multiply
vec2 cmplx_multiply(vec2 a, vec2 b) {
   // so this fits into the distributive property
   // of squaring (x+y)
   // where (x^2 - y^2) is the real number
   // and (xu + yv) is the imaginary
    float real = a.x * b.x - a.y * b.y;
    float imag = a.x * b.y + a.x * b.y;
    return vec2(real, imag);
}

// performs a complex addition operation
vec2 cmplx_add(vec2 a, vec2 b) {
    float real = a.x + b.x;
    float imag = a.y + b.y;
    return vec2(real, imag);
}

// performs a complex subtraction operation
vec2 cmplx_sub(vec2 a, vec2 b) {
    float real = a.x - b.x;
    float imag = a.y - b.y;
    return vec2(real, imag);
}


vec2 cmplx_div(vec2 a, vec2 b) {
    float d = a.x*a.x + a.y * a.y;
    return cmplx_multiply(b, vec2(a.x/d, -(a.y/d)));
}