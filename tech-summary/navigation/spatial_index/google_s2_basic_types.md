# Google S2 Basic Types

## S1Angle

The `S1Angle` class represents a one-dimensional angle.  You could think it equivalent to the value of `latitude` or `longitude`.

```C++
class S1Angle {
  constexpr double radians() const;
  constexpr double degrees() const;

  int32 e5() const;
  int32 e6() const;
  int32 e7() const;


double radians_;
```

`s1angle.h` includes methods for comparison and arithmetic operators, eg:
```C++
// x and y are angles
if (sin(0.5 * x) > (x + y) / (x - y)) { ... }
```

### Radians & Degrees

<img src="../resources/s2_basic_type_what_is_radian.png" alt="s2_basic_type_what_is_radian" width="400"/>

```
circumference = 2π(radius)

1 radian = 180°/π
π radian = 180°
2π radian = 360°
```
<img src="../resources/s2_basic_type_radian_degrees.png" alt="s2_basic_type_radian_degrees" width="400"/>

- [C++ s1angle.h](https://github.com/google/s2geometry/blob/9398b7c8d55c15c4ad7cdc645c482232ea7c087a/src/s2/s1angle.h#L84)
- [Go angle.go](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s1/angle.go#L52)



