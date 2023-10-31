<!--
.. title: #4 Processing tokens and optimizing for fast retrieval
.. slug: notes-on-information-retrieval-2
.. date: 2023-03-16 17:18:00 UTC
.. tags: information retrieval, textbook, tokens, index terms, postings list
.. category: Notes on Information Retrieval
.. link:
.. description: How to go from raw data sequences into index terms and optimizing the postings list data structure for faster retrieval.
.. type: text
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 2: The term vocabulary and postings lists**.

In the previous blog we talked about how to create an inverted index. In this second part, we will look into the intricacies of how a _token_ can be defined and standardized in order to add it into an index as a _term_, and how to use optimize a postings list data structure to ensure faster retrieval.

<br />
## Indexing terms

In very low-level computing terms, every documents must be read into memory as characters and words (see Note 1) with a particular sequence through a decoding algorithm. Depending on the language, there will be different cases on how files are decoded into a sequence of characters and how that language is read. For example, Arabic languages are written from right to left and can include numbers written from left to right. For this section we will work with the Latin alphabet where characters, and more specifically the English language, where words and numbers are written sequentially from left to right.

Once we have a sequence of characters, we can start the task of defining the **document unit** that will be used for indexing. There are some intricacies on how to define _what_ a document unit and a document itself is. For example, emails can contain large threads and attachments - are these all part of the same document? On the other hand, what happens with long documents? Do we take the _entire_ file as our _document_ or do we break it into smaller pieces? Depending on the **granularity** we want in our information retrieval system, there will always be a precision - recall trade off: small documents will be low on precision (many documents might come up), but high on recall and larger documents will be high on recall and low on precision (information will be hard to find).

### Using preprocessing to transform sequences into index-ready terms

After defining our the document unit and decoding our text data into a sequence of characters, we can start working towards creating our **tokens** by breaking down our sequence at specific intervals. In the case of a sentence, tokens can be the words that make up that sentence split by white space. It is important to note that tokens can also be defined as terms or words; however we define tokens as instances that can be grouped together under the same token **type**. The **terms**, on the other hand, are the types that are included in an information retrieval system dictionary, these are related to tokens but include a few preprocessing steps in between, which we will discuss next.

There will always be a challenge defining how to create a tokens. Some words are very straight forward to index, but others take more critical thining. For example, take the word `don't` - shall we tokenize it as `don't`, `dont`, `don` and `t`, or `do` and `n't`? Most of the issues in defining how to tokenize will be language specific so identifying a language before tokenizing is useful to catch and process these language-specific patterns. Whatever strategy we choose we need to ensure that we apply the same strategy at the user retrieval side to ensure the same processing is done and the term types are the same.

Some other issues with tokenization in the English language include:

- How should we tokenize sequences with "special" characters such as proper nouns, urls, monetary values, etc? (e.g. `Turkey` (the country), `https://www.google.com/`, `$100`)
- Should we separate on hyphens (-), when hyphens can be used for naming or grouping words?
- Should we separate on spaces as some sequences don't can mean the wrong thing independently e.g. `New York`, `whitespace` vs `white space`, dates, etc.

After deciding our tokenization strategy, we can start filtering for common words that do not add value to our searches. These words are called **stop words**, the set of stop words is called a **stop list**. Commonly we can group terms by frequency across all collections (**collection frequency**) and filter them our based on a human-in-the-loop assessment, assuming that common words across all collections don't add a lot of information value to any single document. It is important to consider that we should _not always_ discard these stop words as they can be valuable in some phrases (e.g. the song `Let It Be`); however, phrase processing will be discussed later in this blog. Recent trends have also started to keep the stop word list in favour of more robust language processing for indexing and content understanding.

After processing stop words, we can start normalizing our tokens. We do this because there are tokens that mean the same thing but are represented differently such as `UK`, `U.K.`, and `United Kingdom`. **Normalization** is the process of ensuring all our terms are represented equally. The most standard way is to manually create **equivalence classes**, which are dictionaries of "ground truth" terms mapped to different spellings. An alternative is to maintain both normalized and unnormalized terms and use them by keeping an index of unnormalized terms and a dictionary to expand the search query into all unnormalized terms.

There are other processing techniques that can process the differences in how tokens are represented grammatically in language. These are stemming and lemmatization. **Stemming** refers to the process of slicing the ends of the words to remove grammatical differences (e.g. `colored` and `colors` become `color`). Stemming in the English language most commonly relies on `Porter's algorithm` which reduces word sequences in 5 stages. This blog will not go into the details, but these can be found in [the Porter Stemmer official site](http://www.tartarus.org/~martin/PorterStemmer/). On the other hand, **lemmatization** tries to use morphological analysis to analyze the words and transform them into their base dictionary form using techniques from Natural Language Processing (e.g. `saw`, `sees` and `seeing` become `see`). This base form is called the _lemma_. Even though stemming and lemmatization can help in normalizing how we represent tokens, they usually come with marginal gains and other tradeoffs.

<br />
## Updates to the postings list

Just how preprocessing tokens into terms can help in increasing the precision of our information retrieval system, we can also optimize the postings list data structures to make the retrieval computation more efficient and accurate.

### Skip pointers in postings lists

If we want to analyze to two lists of sizes `m` and `n` simultaneously, the time complexity will be `O(m+n)`. To reduce the time complexity below this linear form, we can use a **skip list** - this is a change to the underlying list data structure that contains **skip pointers**. Skip pointers allow us to skip parts of the list that will not appear in the search result. The general idea is that when we step through multiple lists, we can skip to a matched document in another list allowing the document retrieval to be done faster. For example, if we take two posting lists `A: {2,3,4,5,6,7}` and `B:{1,4,5,7}` the moment we match on the document with index `4`, we can skip to the next pointer in both lists.

### Processing phrase queries

It is common that users want to use **phrase queries** such as `New York` and retrieve messages that are familiar with the city and not the words `new` and `york`. Because of this, the reader might be familiar with the use of double quotes (`" "`) around a token to get an exact match. To support these types of phrase queries using exact or close matches, a search system needs to support and properly implement phrase indexes or use proximity indexes that weight terms that appear closer in different documents.

**Biword indexes** or **phrase index** are a type of index that considers words that are together in the document (data practitioners familiar with NLP might relate this to n-grams). Adding these terms into our indexes allows for immediate search amongst them. A downside to this approach is that there might be a large number of false positives. To avoid these we can focus on _noun phrases_ and some preprocessing. For example if we take a noun phrase such as `renegotiation of the constitution` we can perform tokenization and part-of-speech tagging to extract nouns from the sub-tokens and this phrase can become the term `renegotiate constitution`. A phrase index is a good solution but there can be issues resulting in multiple false positives if we combine a phrase index with single word indexes (which is usually required).

For the issues described above, phrase indexes are not the standard solution. Instead, **positional indexes** are most commonly used for IR. To create a positional index, we need to store postings in the form of `docID: {position1, position2, ...}` where each position is the position of the token in the document with the number of times it occurs (expanded into `term, totalFrequency: {docId, frequency:{position1, position2}}`). The retrieval algorithm for phrases will look for documents with both terms and then find documents where the position of the second term is one higher than the first. This retrieval can also be used with `k` positions. Additional notes on sizes of the positional index in Notes 2.

Finally, we can also combine the strategies of phrase indexes and positional indexes. This is done by using phrase indexes for certain queries and positional indexes for other queries. The normal approach is to add common queries in the positional index into the phrase index to ensure faster retrieval.

<br />
## Notes

1.  On document decoding; At the very basic level, every term is a byte in a file. These bytes are decoded into human-readable characters. Usually the document can be decoded using the encoding scheme found in the file's metadata so an in-depth analysis of byte decoding and encoding will are spared in this post. More details on this can be found in the book.
2.  On positional index size; Using positional indexes will increase the size of the stored index as well as the computing complexity.
