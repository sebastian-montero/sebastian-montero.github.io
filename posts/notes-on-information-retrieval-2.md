<!--
.. title: Creating index-ready terms and enhancing the postings list
.. slug: notes-on-information-retrieval-2
.. date: 2023-03-14 23:22:07 UTC
.. tags: information retrieval, textbook, tokens, index terms, postings list
.. category: Notes on Information Retrieval
.. link: 
.. description: Term vocabulary and postings list
.. type: text
.. status: draft
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 2: The term vocabulary and postings lists**.

In the previous blog we talked about how to create an inverted index. In this second part, we will look into the intricacies of how a *term* can be defined and standardized, and how to use optimize a postings list data structure to index these terms.

<br />
## Defining what is a term that should be indexed

Every documents must be read into memory as characters and words (see Notes) with a particular sequence. Depending on the language, there will be edge cases on how files are read for example, Arabic languages are written from right to left and can include numbers written from left to right. For these section we will work with the Latin alphabet where characters, words and numbers are written sequentially from left to right. 

Once a file is read, we can start the task of defining the **document unit** that will be used for indexing. There are some intricacies on how to define *what* a document is. For example, emails can contain large threads and attachments - are these all part of the same document? On the other hand, what happens with long documents? Do we take the *entire* file as our *document* or do we break it into smaller pieces? Depending on the **granularity** we want in our system there will always be a precision - recall trade off: small documents will be low on precision (many documents might come up) but high on recall and larger documents will be high on recall and low on precision (information will be hard to find).

### Using preprocessing to transform sequences into index-ready terms

Once we have a character sequence and a desired document unit, we can start working towards creating our **tokens** by breaking down our sequence. In the case of a sentence these can be the words that make up that sentence. Tokens can also be terms or words; however we define tokens as instances that can be grouped together under the same type. Where a token **type** refers to a class of tokens. The **terms** are the types that are included in an search system dictionary, these are related to tokens but include a few processing steps in between.

There will always be a challenge in the design side when we create tokens. Take the word `don't` - shall we tokenize it as `don't`, `dont`, `don` and `t`, or `do` and `n't`? Most of these issues will be language specific so identifying a language before tokenizing is useful to catch these language-specific patterns. Whatever strategy we choose we need to ensure that we apply the same strategy at the boolean retrieval side to ensure the same processing. 

Some other issues with tokenization in the English language include:
- How should we tokenize sequences with "special" characters such as names, urls, monetary values, etc?
- Should we separate on hyphens (-), when hyphens can be used for naming or grouping words?
- Should we separate on spaces as some sequences don't can mean the wrong thing independently e.g. `New York`, `whitespace` vs `white space`, dates, etc.

After deciding our tokenization strategy, we can start filtering for common words that do not add value to our searches. These sets of words are called **stop words**. Commonly we can group terms by frequency (**collection frequency**) and filter them our based on a human-in-the-loop assessment. It is important to consider that we should *not always* discard these stop words as they can be valuable in some phrases (e.g. the song `Let It Be`). Recent trends have started to keep the stop word list in favour of more robust language processing for indexing and content understanding.

After removing stop words, we can start normalizing our tokens. We do this because there are tokens that mean the same thing but are represented differently such as `UK`, `U.K.`, and `United Kingdom`. **Normalization** is the process of ensuring all our terms are represented equally. The most standard way is to manually create **equivalence classes**, which are dictionaries of "ground truth" terms mapped to different spellings. An alternative is to maintain normalized and unnormalized terms by keeping an index of unnormalized terms and a dictionary to expand the search query into all unnormalized terms. 

Additional processing techniques can process the differences in how tokens are represented in common language: stemming and lemmatization. **Stemming** refers to the process of slicing the ends of the words to remove grammatical differences (e.g. `colored` and `colors` become `color`). Stemming in the English language most commonly relies on `Porter's algorithm` which reduces word sequences in 5 stages. This blog will not go into the details. These can be found [the Porter Stemmer official site](http://www.tartarus.org/~martin/PorterStemmer/). On the other hand, **lemmatization** tries to use morphological analysis to analyze the words and transform them into their base dictionary form - the *lemma* - using techniques from Natural Language Processing. (e.g. `saw`, `sees` and `seeing` become `see`). Even though stemming and lemmatization can help in normalizing how we represent tokens, they usually come with marginal gains and other tradeoffs.


<br />
## Updates to the postings list

Just how we can pre-process tokens into terms to ensure a an accurate retrieval, we can also optimize the postings list data structures to make search more efficient. 

### Skip pointers in postings lists

If we want to analyze to two lists of sizes `m` and `n` simultaneously, the time complexity will be `O(m+n)`. To reduce the time complexity we can use a **skip list** - this is a change to the list that contains **skip pointers**. Skip pointers allow us to skip parts of the list that will not appear in the search result. The general idea is that when we step through multiple lists skip to a matched document in another list.

### Processing phrase queries
 
When we use **phrase queries** such as `New York` we want to ensure that search results match that exact phrase and not the words `new` and `york` independently. Because of this, the use of double quotes (`" "`) around a token in search queries has become a standard syntax in IR. To support these, a search system need to support and properly implement phrase queries and/or use proximity to weight terms that appear closer. 




<br />
## Notes

- At the very basic level, every term is a byte in a file. These bytes are decoded into human-readable characters. Usually the document can be decoded using the encoding scheme found in the file's metadata so an in-depth analysis of byte decoding and encoding will are spared in this post. More details on this can be found in the book. 