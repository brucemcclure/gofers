### How variables are named

![](./docs/variables.png)

### Fundamental data types

![](./docs/fundamental_data_types.png)

### Return value

![](./docs/return_value.png)

### Array vs Slice

An array has a fixed length whereas as a slice has a variable length

Both slices and arrays must be declared with a specific data type. Ie every element in a slice must be of the same type

### Declaring and appending slices

![](./docs/dec_append_slice.png)

### for loops

![](./docs/for_loop_breakdown.png)

- index = index of the element in the list/array
- card = current card we are iterating over
- range cards = Here is a slice of records. This is a key word that will be used to iterate over records.

---

### Cards - OO Approach

In OOP we would create a Deck Class blueprint to create a deck instance. That instance would have some data and functions built into it in order to manipulate the data in the instnace.

![](./docs/oo.png)

Go is 'significantly different'. Instead of classes to create the deck we will be creating a new type inside of Go by extending a base type and adding extra functionality to it. Therefore the data type we will be extending is the slice of strings. The extra functionality we will be adding to the new type will be a function with a receiver, this is a fucntion that belongs to an instance ( Similar to a method on an instance in OOP ).

![](./docs/deck_type.png)
