# Slice Internals

Reference: [Slices in Go: Grow Big or Go Home](https://victoriametrics.com/blog/go-slice/)

## How slice references underlying array

Here is an array that is used as an underlying array.
Just FYI, length and capacity of an array is always same and never changes.

```go
var array = [6]byte{0, 1, 2, 3, 4, 5}

%%

fmt.Printf("array: %p %d %d\n", &array, len(array), cap(array))
```

Slice is merely a reference to an underlying array and it holds an pointer address of the array.  
When a slice starts from the beginning of an array (first element), the address of the slice is exactly same as the address of the array.  
Also the address of index `0` is same.

```go
var slice0 = array[0:3]

%%

fmt.Printf("   slice0: %p %d %d\n", slice0, len(slice0), cap(slice0))
fmt.Printf(" array[0]: %p\n", &array[0])
fmt.Printf("slice0[0]: %p\n", &slice0[0])
```

Slice index simply increment the base pointer address by `index * size of an element`.  
So in the below code, the address of both elements at the same index are same.

```go
%%
for i := range slice0 {
  fmt.Printf("&slice0[%d]: %p == &array[%d]: %p\n", i, &slice0[i], i, &array[i])
}
```

This time there are 2 slices and they start from the second and the third element of the underlying array respectively.  
The address of each slice corresponds to the second and the third element of the array.  

```go
var slice1 = array[1:4]
var slice2 = array[2:5]

%%

fmt.Printf("  slice1: %p %d %d\n", slice1, len(slice1), cap(slice1))
fmt.Printf("array[1]: %p\n", &array[1])
fmt.Printf("  slice2: %p %d %d\n", slice2, len(slice2), cap(slice2))
fmt.Printf("array[2]: %p\n", &array[2])
```

## Slice is a fat pointer

As mentioned earlier, a slice is an reference to an underlying array, bit it is not a single pointer but a **fat** pointer.  
A slice can be casted to a struct and this internal structure (`sliceHeader` struct) is usually referred as the slice header.

```go
type sliceHeader struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	array := [6]byte{0, 1, 2, 3, 4, 5}
	slice := array[1:3]
	header := (*sliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("   array[1]: %p\n", &array[1])
	fmt.Printf("      slice: %p %d %d\n", slice, len(slice), cap(slice))
	fmt.Printf("sliceHeader: %v %d %d\n", header.array, header.len, header.cap)
}
```

## How slice grows

There are 2 slices that point to the same array.  
The third number in the slicing operation (`:4` in `[1:3:4]`) specifies the capasity of the slice as an index of the underlying array.  
It is exclusive so `:4` means up to `array[3]`, and `[1:3:4]` means `array[1]..array[3]`, hence capacity of `slice1` is 3.

```go
var (
    array = [6]byte{0, 1, 2, 3, 4, 5}
    slice1 = array[1:3:4]
    slice2 = array[1:5]
)

%%
fmt.Printf("slice1: %p %d %d\n", slice1, len(slice1), cap(slice1))
fmt.Printf("slice2: %p %d %d\n", slice2, len(slice2), cap(slice2))
```

Now append an element to `slice1`.

```go
%%
slice1 = append(slice1, 6)

fmt.Printf("slice1: %p %d %d\n", slice1, len(slice1), cap(slice1))
fmt.Printf("slice2: %p %d %d\n", slice2, len(slice2), cap(slice2))
fmt.Println(" array:", array)
fmt.Println("slice1:", slice1)
fmt.Println("slice2:", slice2)
```

As you can see, `slice1` still points to the same array and the fourth element of the array was modified.  
This is because append actually modifies the underlying array of appended slice as long as appending does not result in exceeding the capacity of the slice.  
As a consequense, the third element of slice2 was also modified. It seems obvious in hindsight, but it is easy to overlook.

Next, append element multiple times so that it exceeds the capacity of `slice1`.

```go
%%
// Note: the append in the previous cell is no longer effective here.
fmt.Printf("slice1: %p %d %d\n", slice1, len(slice1), cap(slice1))
fmt.Println("slice1:", slice1)

slice1 = append(slice1, 6)
slice1 = append(slice1, 7)

fmt.Printf("slice1: %p %d %d\n", slice1, len(slice1), cap(slice1))
fmt.Printf("slice2: %p %d %d\n", slice2, len(slice2), cap(slice2))
fmt.Println(" array:", array)
fmt.Println("slice1:", slice1)
fmt.Println("slice2:", slice2)
```

The array and `slice2` are modified for the first append, but not the second, because the first append does not exceed the capacity but the second does.  
When appending exceeds the capacity of the appended slice, a new underlying array is created (allocated) from the appended slice and the slice points to the array (so the pointer address of `slice1` differs).

If the capacity is exceeded by a single append operation, the original underlying array is not modified.

```go
%%
// Note: again the append in the previous cell is no longer effective here.
slice1 = append(slice1, 6, 7)

fmt.Printf("slice1: %p %d %d\n", slice1, len(slice1), cap(slice1))
fmt.Printf("slice2: %p %d %d\n", slice2, len(slice2), cap(slice2))
fmt.Println(" array:", array)
fmt.Println("slice1:", slice1)
fmt.Println("slice2:", slice2)
```
