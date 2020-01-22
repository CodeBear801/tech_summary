# Hilbert curve

Hilbert curve helps to downgrade n-dimension data to one dimension and also keeps spatial locality.  

## Examples

2 dimension

<img src="../resources/spatial_index_hilbert_curve_example.png" alt="spatial_index_hilbert_curve_example" width="600"/>

n dimension

<img src="../resources/spatial_index_hilbert_curve_example2.png" alt="spatial_index_hilbert_curve_example2" width="600"/>


## How to generate

### Order 1

<img src="../resources/hilbert_curve_order_1.gif" width="600" height="200" />

- Whether its a `U` or reverse `U` depend on where is `(0,0)`.  In the upper picture, `(0,0)` is set to lower left.

### Order 2

<img src="../resources/hilbert_curve_order_2.gif" width="600" height="200" />

- further dived each cell using order 1's strategy
- connect 4 curves end to end

### Order 3

<img src="../resources/hilbert_curve_order_3.gif" width="600" height="200" />

- Further dived based on order 2
- connect 4 curves end to end

### Order n

<img src="../resources/hilbert_curve_order_n.gif" width="600" height="200" />


## Benefits

### Downgrade n-dimension to 1-dimension

<img src="../resources/hilbert_curve_downgrade.gif" width="400" height="200" />

### Stable result



<img src="../resources/serpentine_shape.gif" width="600" height="200" />

<img src="../resources/hilbert_curve_stable_result.gif" width="600" height="200" />


### Better locality

<img src="../resources/hilbert_curve_locality.png" alt="hilbert_curve_locality" width="400"/>

