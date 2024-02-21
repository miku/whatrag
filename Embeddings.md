# Embeddings

> Then, after decades, embeddings have emerged. We can calculate embeddings for
> words, sentences, and even images. Embeddings are also vectors of numbers,
> but they can capture the meaning.

Text were always vectorized, but based on words, stemmed words, etc. New
elements are a tokenizers, e.g. BPE, that lies between character and word level
and offers a better granularity tradeoff (albeit BPE has been developed in a
complete other context).

The word vectors are then taken e.g. from NN weights for an NN that is trained
to perform a specific task. The vectors then appear to capture meaning, that is
the location of words in that high dimensional space exhibit structure that we
recognize as relationsship on various levels, like grammer, semantics, analogy,
etc.

The embeddings can be calculated very differently, and hence there is
competition about embedding libraries and approaches.

Contrast with classic bag-of-words approach. There, we fix some vocabulary
size, e.g. the number of words or stems we observe or know about. That may be
10K or 100K or the like (470K words in English, overall). It is a sparse space.
High dimension, little data to populate this space. Embeddings have lower
dimensions, maybe 50, 500, or 5000.

The TF-IDF was a step up from basic bag of words model (term frequency, inverse
document freqency). It tries to weight a word by its importance, that is how
common it is. A word like "the" does not contribute much to the overall
meaning, probably.

Representations like "word2vec" (2013) are called dense, because they embed the
words into a lower dimensional space. Thought experiment: imagine a line and
trying to take two words and put them onto the line, e.g. "cat", "car", "dog" -
would you place "cat" closer to "car" or "dog"?

A dimension could be a "semantic axis", e.g. color, word type, frequency,
domain, etc. that is very dependent on the type in input.

Now we have a task to choose, initally there were two tasks:

* CBOW (predict word based on surrounding words)
* skipgram (predict context from word)

In a way, we are trying to compress. We try to learn the weights to map a words
to a lower dimensional space and then have a decoder use that representation to
solve the task.

This was all still based on words as unit, a subword skip-gram approach was
used by "GloVe".

This was an important step.

> For anyone interested in text analysis: PLEASE study and use this code and
> the referenced papers. It's importance is hard to overstate. It is far, far
> better than all previous approaches to word analysis. These representations
> are the dimensional compression that occurs in the middle of a deep neural
> net. The resulting vectors encode rich information about the semantics and
> usage pattern of each word in a very concise way.

> We have barely scratched the surface of the applications of these distributed
> representations. This is a great time to get started in this field - previous
> techniques are almost totally obsoleted by this so everyone is starting from
> the same point. -- [HN6217882](https://news.ycombinator.com/item?id=6217882)

Now you could calculate with words, famously:

```
paris - france + germany ~ berlin
```

That is, the resulting vector from `paris - france + germany` has as its
nearest neighbor (or among them) the vector representing `berlin`.

```
vector('king') - vector('man') + vector('woman') is close to vector('queen')
```

It is like looking into a NN, which often is a black box. It reminds one of the
XXX - search book from oulipo or surrealism?

So this was basically over a decade ago.

> I remember talking about word2vec in an interview in 05/2015, it was still very new then

People were excited.

> But, nothing, absolutely nothing for me has ever come close to what blew my
> mind recently with word2vec: so effortless yet you feel like the model knows
> so much that it has obtained cognitive coherence of the vocabulary. Until
> neuroscientists nail cognition, I am happy to foolishly take that as some
> early form of machine cognition. -- [](https://byterot.blogspot.com/2015/06/five-crazy-abstractions-my-deep-learning-word2doc-model-just-did-NLP-gensim.html)

And following:

> This technique basically trains a model based on a neighborhood window of
> words in a corpus and then projects the result onto [an arbitrary number of]
> n dimensions where each word is a vector in the n dimensional space. Then the
> words can be compared using the cosine similarity of their vectors. And what
> is much more interesting is the arithmetics: vectors can be added or
> subtracted for example vector of Queen is almost equal to King + Woman - Man.
> In other words, if you remove Man from the King and add Woman to it,
> logically you get Queen and but this model is able to represent it
> mathematically.

Perspective of what word embeddings are: [Neural Word Embedding as Implicit Matrix Factorization](https://proceedings.neurips.cc/paper_files/paper/2014/file/feab05aa91085b7a8012516bc3533958-Paper.pdf)

Also: how do you assess how good an embedding is? Cf [Efficient Estimation of Word Representations in Vector Space](https://arxiv.org/pdf/1301.3781.pdf)

Fast forward a few years. The meaning of meaning remains somewhat foggy.

> Given the text "What is the main benefit of voting?", an embedding of the
> sentence could be represented in a vector space, for example, with a list of
> 384 numbers (for example, [0.84, 0.42, ..., 0.02]). Since this list captures
> the meaning, we can do exciting things, like calculating the distance between
> different embeddings to determine how well the meaning of two sentences
> matches. -- [https://huggingface.co/blog/getting-started-with-embeddings](https://huggingface.co/blog/getting-started-with-embeddings)

There are a couple of open source tools for creating embeddings.

> How are embeddings generated? The open-source library called Sentence
> Transformers allows you to create state-of-the-art embeddings from images and
> text for free.



## Readings

* [https://scribe.rip/text-embeddings-comprehensive-guide-afd97fce8fb5](https://scribe.rip/text-embeddings-comprehensive-guide-afd97fce8fb5)
