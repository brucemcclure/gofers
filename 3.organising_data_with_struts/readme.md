### Declaring struts

For now struts can be thought of as `{ hash: ruby, dict: python, object: js }` There is a little more to it than that but that will do for now. The first thing we need to do is to declare the strut, this includes the name of the strut itself, the names of the 'keys' and the type of values on those keys. This will be known as the person strut in this example. Thereafter we can create new values of type person.

![](./docs/strut.png)

![](./docs/basic_strut_code.png)

### Go zero values

Zero values are provided by default when creating a new object of a strut but you dont provide it with any values. ie `var bruce person`. The zero values provided are dependent of the ype of data that the strut holds. As the table below shows. To demonstate this lets print out the empty strut with fmt.Printf("%+v", bruce) This will print out the keys and values.

![](./docs/zero_values.png)

### Updating strut values

Because we now have the default empty values we need to update the values of the strut. This is done in a similar way to other languages. Specify the key and add another vlaue to it.

![](./docs/updating_struts.png)
