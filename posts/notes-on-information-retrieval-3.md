<!--
.. title: Ensuring tolerant retrieval in your search system
.. slug: notes-on-information-retrieval-3
.. date: 2023-03-20 20:54:00 UTC
.. tags: information retrieval, textbook, tolerant retrieval
.. category: Notes on Information Retrieval
.. link: 
.. description: Creating a tolerant retrieval system that supports wildcards, errors and alternate spellings.
.. type: text
.. status: draft
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 3: Dictionaries and tolerant retrieval**. The first section of the chapter is in its separate blog: [How is a dictionary data structure optimized for search?](/posts/notes-on-information-retrieval-3-dictionaries/)

Tolerant retrieval and flexibility at query time is an essential feature for users querying an information retrieval system. This can be achieved by enabling wildcards and building a system that can tolerate spelling errors and alternate spellings.

<br />
## Searching with wildcards

A search system must enable searching with wildcards for several situations:
1. the user is unsure of spelling
2. The user knows of various ways to spell a word and wants an exhaustive result
3. The user wants to capture the stem of the word
4. The user is unsure of a foreign word or phrase

There are different types of wildcard queries. A query such as `eng*` is a **trailing wildcard query** because the `*` symbol occurs at the end of the search term. This type of queries can easily be handled by search tree dictionaries. We can traverse the dictionary until we get to the final letter and then get all the set of `W` items. All of these are then retrieved from our inverted index.

We can also have a **leading wildcard query** where the `*` symbol is at the beginning. For these types of queries we can consider using a **reverse B-tree** in the dictionary, where the root to leaf paths are backwards. 

Using both a types of B-trees we can handle wildcard queries in the middle of a word such as `col*r`. In this case we can use the regular B-tree to find terms starting with `col` and then the reverse B-tree to find the terms ending in `r`. After this we can intersect both lists to find terms that start with `col` and end with `r`. From here we can again use our inverse index to retrieve our desired data.

There are two other techniques to handle wildcard queries. Both express the wildcard query as a boolean query on a special wildcard index. This index considers the boolean query as a superset of terms matching the wildcard index. After we retrieve the terms from the wildcard index, each term is filtered down to the terms matching our wildcard query.

### Permuterm indexes

The **permuterm index** ensures we mark the end of a term with a special character, for example `hello` becomes `hello$`. Then we build the permuterm index by rotating the term and map it to the original term. So a set of `(hello$, ello$h, llo$he, lo$hel, o$hell)` all map to `hello`. The rotated set of indexes is called the **permuterm vocabulary**. This query helps in faster searching because it allows us to rotate the original term with a wildcard to search for the desired terms. For example if we want to search for `he*o`, we then transform the term so the wildcard symbol is at the end so `he*o` becomes `o$he*`. We then look for this term in the permutation index as a leading wildcard query. A permuterm index uses a B-tree to explode its full capabilities, but it is important to note that this type of index could get large very quick. 

### k-gram indexes

To understand **k-gram indexes** we need some initial definitions. First, a **k-gram** is a sequence of `k` characters. We again use a special symbol such as `$` to denote the beginning or end of the term. So in the case of the word `hello` a, 3-gram index would include the words `($he, hel, ell, llo, lo$)`. Then we can index each of these k-grams. 

This type of index then is used to search for each of the k-grams that make up the wildcard term. K-grams sometimes require post-filtering steps that check the original query against the returns terms to filter out terms that don't match. 


<br />
## Spelling errors and alternate spellings



### Edit distance

### K-gram overlaps