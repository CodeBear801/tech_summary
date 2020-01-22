# Google S2 Cell

Give a `(lat, lon)`, how to find cellid for it.  

[Comments from C++ impl -> s2coords.h](https://github.com/google/s2geometry/blob/20c8f339cc9a55fdca1c0e8ab519da399752e70b/src/s2/s2coords.h#L38)
```C++
//  (id)
//    An S2CellId is a 64-bit encoding of a face and a Hilbert curve position
//    on that face.  The Hilbert curve position implicitly encodes both the
//    position of a cell and its subdivision level (see s2cell_id.h).
```
You could find similar comments in [go's impl -> stuv.go](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/stuv.go#L35)  

## Step 1: p=(lat, lon) -> (x, y, z)

How lat, lon is generated for a point:  
<img src="../resources/google_s2_world_lat_lon.png" alt="google_s2_world_lat_lon" width="400"/>

[Code from go impl](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/latlng.go#L85)
```go
func PointFromLatLng(ll LatLng) Point {
	phi := ll.Lat.Radians()
	theta := ll.Lng.Radians()
	cosphi := math.Cos(phi)
	return Point{r3.Vector{math.Cos(theta) * cosphi, math.Sin(theta) * cosphi, math.Sin(phi)}}
}
```

<img src="../resources/google_s2_lat_lon_xyz.png" alt="google_s2_lat_lon_xyz" width="400"/>

(Ref from [Spherical to Cartesian coordinates Calculator](https://keisan.casio.com/exec/system/1359534351))

`θ`即为经纬度的纬度，也就是上面代码中的`phi`，`φ`即为经纬度的经度，也就是上面代码的`theta`
```
x = r * cos θ * cos φ
y = r * cos θ * sin φ 
z = r * sin θ
```

[Comments from C++ impl -> s2coords.h](https://github.com/google/s2geometry/blob/20c8f339cc9a55fdca1c0e8ab519da399752e70b/src/s2/s2coords.h#L38)
```C++
//  (x, y, z)
//    Direction vector (Point). Direction vectors are not necessarily unit
//    length, and are often chosen to be points on the biunit cube
//    [-1,+1]x[-1,+1]x[-1,+1]. They can be be normalized to obtain the
//    corresponding point on the unit sphere.
//
//  (lat, lng)
//    Latitude and longitude (LatLng). Latitudes must be between -90 and
//    90 degrees inclusive, and longitudes must be between -180 and 180
//    degrees inclusive.
```

[Comments from s2cell_hierarchy](https://s2geometry.io/devguide/s2cell_hierarchy.html)

```
(x, y, z)\ Spherical point: The final S2Point is obtained by projecting the (face, u, v) coordinates onto the unit sphere. Cells in (x,y,z)-coordinates are quadrilaterals bounded by four spherical geodesic edges.
```


## Step 2: (x, y, z) -> (face, u, v)

<img src="../resources/google_s2_cellid_step2.png" alt="google_s2_cellid_step2" width="400"/>

[code](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/stuv.go#L229)
```go
// xyzToFaceUV converts a direction vector (not necessarily unit length) to
// (face, u, v) coordinates.
func xyzToFaceUV(r r3.Vector) (f int, u, v float64) {
	f = face(r)
	u, v = validFaceXYZToUV(f, r)
	return f, u, v
}
```

**其实就是在圆心打出灯光，把地球表面投影到立方体的一个面上**

<img src="../resources/google_s2_cellid_step2_world2face.png" alt="google_s2_cellid_step2_world2face" width="400"/>

(face,u,v) 表示一个立方空间坐标系，三个轴的值域都是 `[-1,1]` 之间。


[Comments from s2cell_hierarchy](https://s2geometry.io/devguide/s2cell_hierarchy.html)

```
(face, u, v)\ Cube-space coordinates: To make the cells at each level more uniform in size after they are projected onto the sphere, we apply a nonlinear transformation of the form u=f(s), v=f(t) before projecting points onto the sphere. This function also scales the (u,v)-coordinates so that each face covers the biunit square [-1,1]×[-1,1]. Cells in (u,v)-coordinates are rectangular, and are not necessarily subdivided around their center point (because of the nonlinear transformation “f”).
```

## Step 3: Non-linear transform (face, u, v) -> (face, s, t)

<img src="../resources/google_s2_cellid_step3.png" alt="google_s2_cellid_step3" width="400"/>



### Problem

投影比例不同的问题

google_s2_cellid_step3_projection_Issue

在两级的投影比在赤道区域要大.  需要一个转换将上面长的拉短、将下面短的拉长，尽量让区间变相同。
Similar issue also exists in mercator projection, you could find this interesting article from CNN: [What's the real size of Africa? How Western states used maps to downplay size of continent](https://www.cnn.com/2016/08/18/africa/real-size-of-africa/index.html)

### Solutions

google_s2_cellid_step3_trade_off

```
// linear
u = 0.5 * ( u + 1)

// tan() 
u = 2 / pi * (atan(u) + pi / 4) = 2 * atan(u) / pi + 0.5

// Quadratic
u >= 0，u = 0.5 * sqrt(1 + 3*u)
u < 0, u = 1 - 0.5 * sqrt(1 - 3*u)
```

```
线性变换是最快的变换，但是变换比最小。
tan() 变换可以使每个投影以后的矩形的面积更加一致，最大和最小的矩形比例仅仅只差0.414。
但是 tan() 函数的调用时间非常长。如果把所有点都按照这种方式计算的话，性能将会降低3倍。
最后谷歌选择的是二次变换，这是一个近似切线的投影曲线。它的计算速度远远快于 tan() ，
大概是 tan() 计算的3倍速度。生成的投影以后的矩形大小也类似。
不过最大的矩形和最小的矩形相比依旧有2.082的比率。
```

[code](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/stuv.go#L176)

```go
// stToUV converts an s or t value to the corresponding u or v value.
// This is a non-linear transformation from [-1,1] to [-1,1] that
// attempts to make the cell sizes more uniform.
// This uses what the C++ version calls 'the quadratic transform'.
func stToUV(s float64) float64 {
	if s >= 0.5 {
		return (1 / 3.) * (4*s*s - 1)
	}
	return (1 / 3.) * (1 - 4*(1-s)*(1-s))
}

// uvToST is the inverse of the stToUV transformation. Note that it
// is not always true that uvToST(stToUV(x)) == x due to numerical
// errors.
func uvToST(u float64) float64 {
	if u >= 0 {
		return 0.5 * math.Sqrt(1+3*u)
	}
	return 1 - 0.5*math.Sqrt(1-3*u)
}
```

[Comments from s2cell_hierarchy](https://s2geometry.io/devguide/s2cell_hierarchy.html)

```
(face, s, t)\ Cell-space coordinates: “s” and “t” are real numbers in the range [0,1] that identify a point on the given face. For example, the point (s, t) = (0.5, 0.5) corresponds to the center of the cell at level 0. Cells in (s, t)-coordinates are perfectly square and subdivided around their center point, just like the Hilbert curve construction.
```
**u，v的值域是`[-1,1]`，变换以后，是s，t的值域是`[0,1]`。**


## Step 4: (face, s, t) -> (face, i, j)


<img src="../resources/google_s2_cellid_step4.png" alt="google_s2_cellid_step4" width="400"/>


**s，t的值域是`[0,1]`，现在值域要扩大到`[0,2^30^-1]`。**

[code](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/cellid.go#L635)
```go
// stToIJ converts value in ST coordinates to a value in IJ coordinates.
func stToIJ(s float64) int {
	return clampInt(int(math.Floor(maxSize*s)), 0, maxSize-1)
}

// clampInt returns the number closest to x within the range min..max.
func clampInt(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// caller
// cellIDFromFaceIJ(f, stToIJ(0.5*(u+1)), stToIJ(0.5*(v+1)))
```

[Comments from s2cell_hierarchy](https://s2geometry.io/devguide/s2cell_hierarchy.html)
```
(face, i, j)\ Leaf-cell coordinates: The leaf cells are the subsquares that result after 30 levels of Hilbert curve subdivision, consisting of a 230 × 230 array on each face. “i” and “j” are integers in the range [0, 230-1] that identify a particular leaf cell. The (i, j) coordinate system is right-handed on every face, and the faces are oriented such that Hilbert curves connect continuously from one face to the next.

```

## Step 5: Generate Cell id
[Comments from s2cell_hierarchy](https://s2geometry.io/devguide/s2cell_hierarchy.html)
```
Cell id: A 64-bit encoding of a face and a Hilbert curve parameter on that face, as discussed above. The Hilbert curve parameter implicitly encodes both the position of a cell and its subdivision level.
```
<img src="../resources/google_s2_cellid_step5_1.png" alt="google_s2_cellid_step5_1" width="400"/>

<img src="../resources/google_s2_cellid_step5_2.png" alt="google_s2_cellid_step5_2" width="400"/>

<img src="../resources/google_s2_cellid_step5_3.png" alt="google_s2_cellid_step5_3" width="400"/>

You could think the problem as, for given `(i, j)`, how could you find corresponding s2 cell.  

There are only three part of code matters from [cellid.go](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/cellid.g0)

1. Definition

```go
// Constants related to the bit mangling in the Cell ID.
const (
	lookupBits = 4
	swapMask   = 0x01
	invertMask = 0x02
)

// The following lookup tables are used to convert efficiently between an
// (i,j) cell index and the corresponding position along the Hilbert curve.
//
// lookupPos maps 4 bits of "i", 4 bits of "j", and 2 bits representing the
// orientation of the current cell into 8 bits representing the order in which
// that subcell is visited by the Hilbert curve, plus 2 bits indicating the
// new orientation of the Hilbert curve within that subcell. (Cell
// orientations are represented as combination of swapMask and invertMask.)
//
// lookupIJ is an inverted table used for mapping in the opposite
// direction.
//
// We also experimented with looking up 16 bits at a time (14 bits of position
// plus 2 of orientation) but found that smaller lookup tables gave better
// performance. (2KB fits easily in the primary cache.)
var (
	ijToPos = [4][4]int{
		{0, 1, 3, 2}, // canonical order
		{0, 3, 1, 2}, // axes swapped
		{2, 3, 1, 0}, // bits inverted
		{2, 1, 3, 0}, // swapped & inverted
	}
	posToIJ = [4][4]int{
		{0, 1, 3, 2}, // canonical order:    (0,0), (0,1), (1,1), (1,0)
		{0, 2, 3, 1}, // axes swapped:       (0,0), (1,0), (1,1), (0,1)
		{3, 2, 0, 1}, // bits inverted:      (1,1), (1,0), (0,0), (0,1)
		{3, 1, 0, 2}, // swapped & inverted: (1,1), (0,1), (0,0), (1,0)
	}
	posToOrientation = [4]int{swapMask, 0, 0, invertMask | swapMask}
	lookupIJ         [1 << (2*lookupBits + 2)]int
	lookupPos        [1 << (2*lookupBits + 2)]int
)
```

2. [Init](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/cellid.go#L721)
```go
func init() {
	initLookupCell(0, 0, 0, 0, 0, 0)
	initLookupCell(0, 0, 0, swapMask, 0, swapMask)
	initLookupCell(0, 0, 0, invertMask, 0, invertMask)
	initLookupCell(0, 0, 0, swapMask|invertMask, 0, swapMask|invertMask)
}

// initLookupCell initializes the lookupIJ table at init time.
func initLookupCell(level, i, j, origOrientation, pos, orientation int) {
	if level == lookupBits {
		ij := (i << lookupBits) + j
		lookupPos[(ij<<2)+origOrientation] = (pos << 2) + orientation
		lookupIJ[(pos<<2)+origOrientation] = (ij << 2) + orientation
		return
	}

	level++
	i <<= 1
	j <<= 1
	pos <<= 2
	r := posToIJ[orientation]
	initLookupCell(level, i+(r[0]>>1), j+(r[0]&1), origOrientation, pos, orientation^posToOrientation[0])
	initLookupCell(level, i+(r[1]>>1), j+(r[1]&1), origOrientation, pos+1, orientation^posToOrientation[1])
	initLookupCell(level, i+(r[2]>>1), j+(r[2]&1), origOrientation, pos+2, orientation^posToOrientation[2])
	initLookupCell(level, i+(r[3]>>1), j+(r[3]&1), origOrientation, pos+3, orientation^posToOrientation[3])
}

```

3. [Query](https://github.com/golang/geo/blob/5b978397cfecc7280e598e9ac5854e9534b0918b/s2/cellid.go#L564)
```go
// cellIDFromFaceIJ returns a leaf cell given its cube face (range 0..5) and IJ coordinates.
func cellIDFromFaceIJ(f, i, j int) CellID {
	// Note that this value gets shifted one bit to the left at the end
	// of the function.
	n := uint64(f) << (posBits - 1)
	// Alternating faces have opposite Hilbert curve orientations; this
	// is necessary in order for all faces to have a right-handed
	// coordinate system.
	bits := f & swapMask
	// Each iteration maps 4 bits of "i" and "j" into 8 bits of the Hilbert
	// curve position.  The lookup table transforms a 10-bit key of the form
	// "iiiijjjjoo" to a 10-bit value of the form "ppppppppoo", where the
	// letters [ijpo] denote bits of "i", "j", Hilbert curve position, and
	// Hilbert curve orientation respectively.
	for k := 7; k >= 0; k-- {
		mask := (1 << lookupBits) - 1
		bits += ((i >> uint(k*lookupBits)) & mask) << (lookupBits + 2)
		bits += ((j >> uint(k*lookupBits)) & mask) << 2
		bits = lookupPos[bits]
		n |= uint64(bits>>2) << (uint(k) * 2 * lookupBits)
		bits &= (swapMask | invertMask)
	}
	return CellID(n*2 + 1)
}
```

