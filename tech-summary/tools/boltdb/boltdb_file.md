

## Summary

One db will be record in single file on disk, which contains a list of `pages`    
<img src="https://user-images.githubusercontent.com/16873751/98424592-f844f380-2046-11eb-99f4-021003288f93.png" alt="botdb_bucket" width="600"/>
<br/>


When load page into **memory**, there is conversion from [page -> node](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/node.go#L191), then then node is organized as two dimension tree.  Meta is a spacial page as entrance to all data to certain db.  
<img src="https://user-images.githubusercontent.com/16873751/98425036-98e7e300-2048-11eb-9105-3e428d00d256.png" alt="botdb_bucket" width="600"/>
<br/>

When operate on the tree, read or write, an explore process which is abstract as `bucket`, with the help of `cursor` to find specific location
<img src="https://user-images.githubusercontent.com/16873751/98408842-aee4ac00-2026-11eb-9d67-d807ddb8104a.png" alt="botdb_bucket" width="600"/>
<br/>

## db
Please go to [boltdb_internal->db](./boltdb_internal.md/#db) for more information.


## Page
[page's definition](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/page.go#L30)

```go
type page struct {
	id       pgid
	flags    uint16  // [perry] there are 4 kinds of page
	count    uint16  // [perry] only valid when page is branch or leaf
	overflow uint32  // [perry] when data is too large not fit into a single page, records additional page number needed
	ptr      uintptr  //[perry] start location of page data
}
```

About `ptr`

<img src="https://user-images.githubusercontent.com/16873751/98424599-fed36b00-2046-11eb-9966-de2be3f633e6.png" alt="botdb_bucket" width="200"/>
<br/>

4 kinds of page

```go
const (
	branchPageFlag   = 0x01
	leafPageFlag     = 0x02
	metaPageFlag     = 0x04
	freelistPageFlag = 0x10
)

```

### Write page

[page->hexdump()](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/page.go#L85) converts structure directly to byte array and write into file

```go
// dump writes n bytes of the page to STDERR as hex output.
func (p *page) hexdump(n int) {
	buf := (*[maxAllocSize]byte)(unsafe.Pointer(p))[:n]
	fmt.Fprintf(os.Stderr, "%x\n", buf)
}
```

### Read page

[db->page(pageid)](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/db.go#L792) loads page from pageid

```go
// page retrieves a page reference from the mmap based on the current page size.
func (db *DB) page(id pgid) *page {
	pos := id * pgid(db.pageSize)
	return (*page)(unsafe.Pointer(&db.data[pos]))
}
```

## meta page

<img src="https://user-images.githubusercontent.com/16873751/98425887-b23e5e80-204b-11eb-97e6-f4298e7c6e2c.png" alt="botdb_bucket" width="600"/>
<br/>


## freelist


<img src="https://user-images.githubusercontent.com/16873751/98425933-e023a300-204b-11eb-8c3c-4ea6ad37c09f.png" alt="botdb_bucket" width="600"/>
<br/>

## Branchpage

`Branchpage` records index data of B+ tree, it will be loaded as `Branch Node`

<img src="https://user-images.githubusercontent.com/16873751/98426117-7b1c7d00-204c-11eb-88cf-a11cc531ddd5.png" alt="botdb_bucket" width="600"/>
<br/>

[branch page operations](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/page.go#L71)
```go

// branchPageElement retrieves the branch node by index
func (p *page) branchPageElement(index uint16) *branchPageElement {
	return &((*[0x7FFFFFF]branchPageElement)(unsafe.Pointer(&p.ptr)))[index]
}

// branchPageElements retrieves a list of branch nodes.
func (p *page) branchPageElements() []branchPageElement {
	if p.count == 0 {
		return nil
	}
	return ((*[0x7FFFFFF]branchPageElement)(unsafe.Pointer(&p.ptr)))[:]
}
```

