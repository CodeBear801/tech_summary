# mongodb



Start

```
bin/mongod --directoryperdb --dbpath mongodb/data/db --logpath mongodb/log/mongo.log --logappend --rest --install
```

DB operations

```
// list existing db
show dbs             
// create a new db          
use mycustomers        
// current db
db                                  
```

Collections
```
// create collections
db.createCollections   
// show all collections
show collections        
// https://www.tutorialspoint.com/mongodb/mongodb_create_collection.htm
```

Operations

```
db.posts.insertMany([{}, {}, {}]) 

{
	title: 'Part two',
	body: 'Body of post two',
	category: 'Techology',
	date: Date()
}

db.posts.find()          
// find elements

db.posts.find({category: 'Techology' }).pretty()

db.posts.find({category: 'Techology' }).count()

db.posts.find().limit(2)

db.posts.find().sort({ title: 1}).pretty()

db.posts.find().sort({ title: -1}).pretty()

db.posts.find().forEach(function(doc) { print ('Blog Post: ' + doc.title ) })
// JS function

db.posts.update({title: 'Post Two'}, { title: 'Post Two', body: 'New post 2 body', date: Date()}, { upsert:true })
// will use new one replace old one

db.posts.update({title: 'Post Two'}, { $set: { body: 'New post 2 body2' } } )
// keep original, only replace new colom

db.posts.update({title: 'Post Two'}, { $inc: { likes: 2 } } )

db.posts.update({title: 'Post Two'}, { $rename: { likes: 'views' } } )
// change name from 'likes' to 'views'

db.posts.remove({title: 'Post Two'})

db.posts.update({title: 'Post Two'}, { $set: { comments: [{user: 'lx', body: 'I like it', date: Date()}, {user: 'lx', body: 'I like it', date: Date()} ] } } )
// add new array

db.posts.find({ comments:{ $elemMatch: {user: 'lx'} }     })
// search data in array

db.posts.createIndex({ title : 'text' })
// create index for title

db.posts.find({ $text: { $search: "\"Post O\""}})
// text search

db.posts.update({title: 'Post Two'}, { $set: {views: 10}})
db.posts.find({})
db.posts.find( {views: { $gt:3 } } )
// gt, gte, lt, lte

```


Reference 
- [MongoDB Crash Course 2019](https://www.youtube.com/watch?v=-56x56UppqQ)




