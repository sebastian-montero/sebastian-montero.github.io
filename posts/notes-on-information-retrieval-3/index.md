<!--
.. title: Notes on IR: Tolerant retrieval with text
.. slug: notes-on-information-retrieval-3
.. date: 2023-03-21 00:26:00 UTC
.. tags: information retrieval, textbook, tolerant retrieval
.. category: Notes on Information Retrieval
.. link:
.. description: Creating a tolerant retrieval system that supports wildcards, errors and alternate spellings.
.. type: text
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 3: Dictionaries and tolerant retrieval**. The first section of the chapter is in its separate blog: [How is a dictionary data structure optimized for search?](/posts/notes-on-information-retrieval-3-dictionaries/)

Tolerant retrieval is an essential for users querying an information retrieval system. Tolerance at query time includes fixing user mistakes such as spelling mistakes, alternative spelling, phonetic corrections and also includes the flexibility for users to use wildcards during their search.

This section will introduce wildcards and query correction.

<br />
## Searching with wildcards

Many search systems enable searching with wildcards, so it has become a standard in modern systems. There are many reasons why a user might want to use wildcards. These include:

1. The user is unsure of the right spelling for a term
2. The user knows of various ways to spell a word and wants an exhaustive result (e.g. `col*r` for `color` and `colour`)
3. The user wants to capture the stem of the word (e.g. `fl*` for `fly` and `flew`)
4. The user is unsure of a foreign word or phrase

In this section we will introduce are different types of wildcard queries. A query where the wildcard symbol `*` (e.g. `eng*`) occurs at the end of the term is known as a **trailing wildcard query**. This type of query can easily be handled by search tree dictionaries - specifically B-trees. With a B-tree, we can traverse the dictionary until we get to the final letter and then get all the terms under that final letter. All of the documents matching these terms are then retrieved from our inverted index. We can also have a **leading wildcard query** where the wildcard symbol `*` is at the beginning fo the term. For these types of queries we can consider using a **reverse B-tree** as our dictionary, where the root to leaf paths are backwards.

Using both types of B-trees we can handle wildcard queries in the middle of a word such as `col*r`. In this case we can use the regular B-tree to find terms starting with `col` and then the reverse B-tree to find the terms ending in `r`. After we traverse both search trees, we can intersect both sets of results to find terms that start with `col` and end with `r`. From here we can again use our inverse index to retrieve target documents with these terms.

Now we will introduce two additional wildcard query search algorithms that express the wildcard query as a boolean query on a _special wildcard index_. This specialized index considers the boolean query as a superset of terms that match the wildcard index. After we retrieve the terms from the wildcard index, each term is filtered down to the terms matching our query.

### Permuterm indexes

First, we introduce the **permuterm index**. This index uses a special character to match the end of each term that will be indexed (for example `hello` becomes `hello$`). Then we build the permuterm index by rotating the term and map it to the original term. So a set of `(hello$, ello$h, llo$he, lo$hel, o$hell)` all map to `hello`. The rotated superset of terms is called the **permuterm vocabulary**. This query helps in faster searching because it allows us to rotate the original term with a wildcard to search for the desired term permutations. If we want to search for `he*o`, we then transform the term so the wildcard symbol is at the end so `he*o` becomes `o$he*`. We then look for the `o$he` term in the permutation index as a leading wildcard query. A permuterm index uses a B-tree. It is important to note that this type of index could get large very quick.

### k-gram indexes

The next type of index is a **k-gram index**. To build this specialised wildcard index, we need to build a set of k-grams for each term. a **k-gram** is a sequence of `k` characters that are adjacent in a term. We again use a special symbol such as `$` to denote the beginning and the end of the term. So in the case of the word `hello`, a 3-gram index would include the words `($he, hel, ell, llo, lo$)`. Then we can index each of these k-grams in our k-gram index. When we search for `he*o`, we then search for the k-grams that make up this phrase. K-grams sometimes require post-filtering steps that check the original query against the returns terms to filter out terms that don't match.

<br />
## Spelling errors, alternate spellings and phonetic corrections

When it comes to spelling corrections and alternate spellings, there are two principles that need to be followed: First, the system should always chose the nearest term based on a number of alternative correct spellings. This requires that our system has an understanding of term proximity. Second, if you have more than one match to a correct term, select the most common one. You can do this by looking at the occurrences of the term in the collection.

Before we will introduce spelling correction algorithms, we can discuss how we run these corrections. We can use **isolated-term** correction, this is the type of correction that will identify a misspelled term and attempt to correct each term individually. Note that the system will not understand if the query as a whole makes sense, just its individual parts, hence the phrase `flew form Heathrow` would be correct as each term is individually correct. On the other hand, we also have **context-sensitive** correction which considers the search phrase as a whole. We will proceed to look into two algorithms for isolated-term corrections: edit distance and k-gram overlap.

The **edit distance** algorithm takes two strings `s1` and `s2` and defines the distance as the number of edit operations required to transform `s1` to `s2`. The edit operations can include: inserting a character, deleting a character, replacing a character. The edit distance can also be called the **Levenshtein distance**. The edit distance can also be changed to allow for different weights for different operations. Setting weights for the likelihood of making a change or a mistake is very effective. I will not go into detail of the algorithm in this post; however this can be found in section 3.3.3 of the book.

Another method for spelling correction using k-gram indexes is called **k-gram overlap**. We can use k-gram indexes to limit the set of vocabulary terms for which we compute the edit distances. Using a k-gram index will results in many terms that have common k-grams. From here we will need to do a linear scan through the postings to filter for the k-grams that overlap with the search query.

In a retrieval system, we might want to first look into isolated-term correction if it identifies an issue with the query, but if there are a small number of documents returned without a any correction or after isolated-term corrections, we might want the system to do context-sensitive correction. To do this we need to provide alternative spellings for all the terms being searched and try to substitute them in the query. This operation can be computationally expensive so we could use simple heuristics like looking at the most frequent alternative spelling only or look at bi-word statistics and find what words are usually used together.

Another form of correction is the phonetic corrections. This type of correction requires a specialized index and which considers how terms sound. To build a phonetic search index, we need to create a `phonetic hash` that map similar word sounds to the same hash. Phonetic hashing is collectively known as **soundex** algorithms. At the most basic level a soundex algorithm will turn every indexable term into a hash form and build a soundex index from that, then do the same for query terms and finally use the soundex index to see if there is a match to the soundex query. Even though we are not going too much in the detail of soundex algorithms, they work very well, specially with names. A caveat of this indexes, however, is that they are alphabet dependent.
