## streetaddresssimkey-go

Go package for generating a similarity key from the input data used to match with other similar street address data. Use the generated similarity key, rather than the actual data itself, to match and/or sort street address data by similarity. This avoids the problems of data inconsistency, misspellings, and address variations when matching within a single dataset, and can also help matching across datasets or for more advanced searching.

### Usage

To generate the similarity key, you will need the following information:

- an API License key, available at https://www.interzoid.com
- a street address to generate the similarity key for

Begin by importing the package:

    import "github.com/interzoid/streetaddresssimkey-go"

Then, feed the information into the GetSimKey() method:

    score, code, credits, err := StreetAddressSimKey.GetSimKey("YOUR-API-KEY","107 E. Main St.")


The return values will be the similarity key, a code (success or failure), how many remaining credits on your API key, and any error messages. The similarity key can be used to search for other similar street addresses, to sort large datasets by similarity, and perhaps use additional attributes to identify duplicates/redundancy.