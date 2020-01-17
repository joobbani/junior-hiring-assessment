# Junior Hiring Assessment

### 🎉 Thank you for applying to PDW and congratulations on making it to this step. 🎉

As part of your application process, this assignment is your opportunity to provide us with answers to two important questions here at PDW:

- How well can you learn Go?
- What level of domain knowledge do you already possess?

</br>

### Within the following _contraints_, create a JSON API that fulfills all _requirements_ below:

#### Constraints:

- Use the included `person.go` file within your code base exactly as provided
- Only utilize the standard library; meaning no third party libraries such as `gorilla/mux` or `valyala/fastjson`

#### Requirements:

Create a JSON API that will:

- Utilize a map to store the people
- Consume `Person` structs via a **POST** request to /people
- Retrieve and return the marshaled JSON for the appropriate person with a **GET** request to /people/{name}
- Write out to a file _AND_ return the marshaled json for _all people_ in response to a **GET** request to /people
