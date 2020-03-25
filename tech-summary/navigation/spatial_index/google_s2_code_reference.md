# Google S2 code references

## sidewalklabs
- Regioncoverer URL: http://s2.sidewalklabs.com/regioncoverer/
- http://s2.sidewalklabs.com/static/regioncoverer.js?v=0.3.0
- https://github.com/sidewalklabs/s2sphere/blob/master/s2sphere/sphere.py

## s2demo
- URL: http://s2map.com
- http://s2map.com/llmap.js

```js
/**
 * @param {Array.<fourSq.api.models.geo.S2Response>} cells
 * @return {Array.<L.Polygon>}
 */
renderCellsForHeatmap: function(cellColorMap, cellDescMap, cellOpacityMap, cells) {
```

- convert circle to polygon http://s2map.com/thirdparty/leaflet-geodesy.js


```js

module.exports.circle = function(center, radius, opt) {
    center = L.latLng(center);
    opt = opt || {};
    var parts = opt.parts || 20;

    function generate(center) {
        var lls = [];
        for (var i = 0; i < parts + 1; i++) {
            lls.push(spherical.radial(
                [center.lng, center.lat],
                (i / parts) * 360, radius).reverse());
        }
        return lls;
    }

    var poly = L.polygon(generate(center), opt);

    poly.setLatLng = function(_) {
        center = _;
        poly.setLatLngs(generate(center));
        return poly;
    };

    poly.getRadius = function(_) {
        return radius;
    };

    poly.setRadius = function(_) {
        radius = _;
        poly.setLatLngs(generate(center));
        return poly;
    };

    return poly;
};
```

https://github.com/gabzim/circle-to-polygon  
https://github.com/gabzim/circle-to-polygon/blob/master/index.js  





