- [Why column data](#why-column-data)
- [Understanding the case in paper](#understanding-the-case-in-paper)

more info: https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/navigation/eta/team_discussion_for_06182020.md#why-parquet


## Why column data
Target of column based data organize
- Only load needed data
- Organize data storage based on needs(query, filter)

For flat data

id|column_1|column_2|column_3
---|---|---|---
1|111|222|333
2|-|-|444
3|555|666|777

Easy to encoding and decoding.  For `none-exist` field, we could use some `marker field` to record it.

But when we talk about `nested data`, take `protobuf` as example, it has
```
required: exactly one occurrence
optional: 0 or 1 occurrence
repeated: 0 or more occurrences
```
For example
```
message AddressBook {
   required string owner;
   repeated string ownerPhoneNumbers;
   repeated group contacts {
        required string name;
        optional string phoneNumber;
   }
}

```
How can we re-construct original structure with just value(plus some `marker`)

## Understanding the case in paper

`Definition Level`  
To support nested records we need to store the level for which the field is `null`. This is what the definition level is for: from 0 at the root of the schema up to the maximum level for this column. When a field is defined then all its parents are defined too, but when it is `null` we need to record the level at which it started being null to be able to reconstruct the record.


`Repetition levels`  
To support repeated fields we need to store when new lists are starting in a column of values. This is what repetition level is for: it is the level at which we have to create a new list for the current value. In other words, the repetition level can be seen as a marker of when to start a new list and at which level. 

<img src="https://user-images.githubusercontent.com/16873751/85639756-9c04a880-b63e-11ea-9d47-ac10d09e1d0b.png" alt="dremel_pic1" width="500"/>  <br/>

<img src="https://user-images.githubusercontent.com/16873751/85639769-a2932000-b63e-11ea-979e-36e3521b1d58.png" alt="dremel_pic1" width="500"/>  <br/>


<img src="https://user-images.githubusercontent.com/16873751/85639854-e71ebb80-b63e-11ea-8f15-6bbb4da54238.png" alt="dremel_pic1" width="800"/>  <br/>

<img src="https://user-images.githubusercontent.com/16873751/85639859-ebe36f80-b63e-11ea-8095-c9bf68a4cfe4.png" alt="dremel_pic1" width="500"/>  <br/>

create one column per primitive type cell shown in blue



Column | Max Definition Level |Reason| Max Repetition Level| Reason
---|---|---|---|---
DocID|0|required|0|no repetition
forward|2||1|
backward|2||1|
url|2||1|
code|2|required, use its parent level, or understand as when there is `language` there must exists `code`|2|
Country|3||2|


Explanation for different value of `Repetition level`, take `Country` as example
- 0 marks every new record and implies creating a new `Name` and `Language` List
- 1 marks every new record and implies creating a new `Language` List
- 2 marks every new element in Language


Explanation of `NULL`
- In `Name.Url`, totally there are **4** names in data, r = 1 means a new entry in the Name at Level 1, d = 1 means Name is defined but not URL, so just create an empty URL
- In `Links.Backward`, r = 0 means a new record and creates new list, d = 1 means `Links` is defined but not `Backward`
- For `Code` and `Country`, there are **5** field has language due to that first name has **2** language.
