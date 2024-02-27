# Ouroboros

> [Ouroboros](https://en.wikipedia.org/wiki/Ouroboros), or how to build a
> retrieval-augmented data generative AI pipeline from scratch

![](static/Chrysopoea.png)

A large language model can be flaky
[compressors](https://arxiv.org/pdf/2309.10668.pdf). How to improve reliability
of the output? One approach is to retrieval-augmented generation (Meta, 2020).

> Underpinning all foundation models, including LLMs, is an AI architecture
> known as the **transformer**. It turns heaps of raw data into a compressed
> representation of its basic structure. Starting from this raw representation,
> a foundation model can be adapted to a variety of tasks with some additional
> fine-tuning on labeled, domain-specific knowledge.

> But fine-tuning alone rarely gives the model the full breadth of knowledge it
> needs to answer highly specific questions in an ever-changing context. In a
> 2020 paper, Meta came up with a framework called retrieval-augmented
> generation to give LLMs access to information beyond their training data. RAG
> allows LLMs to build on a specialized body of knowledge to answer questions
> in more accurate way. -- [What is retrieval-augmented generation?](https://research.ibm.com/blog/retrieval-augmented-generation-RAG)

Can this approach help finding code, documentation, text passages faster?

> Countless hours spent searching for documentation on peculiar, intellectually
> uninteresting aspects; the efforts to learn an overly complicated API, often
> without good reason; writing immediately usable programs that I would discard
> after a few hours. These are all things I do not want to do, especially now,
> with Google having become a sea of spam in which to hunt for a few useful
> things. -- [http://antirez.com/news/140](http://antirez.com/news/140)

and also:

> I have also learned that LLMs are a bit like Wikipedia and all the video
> courses scattered on YouTube: **they help those with the will, ability, and
> discipline**, but they are of marginal benefit to those who have fallen behind.

How do we learn, when and how to use this technology?

> Have a peek at this blog post that is going around lately: [The pain points
> of building a copilot](https://austinhenley.com/blog/copilotpainpoints.html)
> These people are brimming with excitement about all the new problems that
> LLMs are bringing to the table. -- [Why We Can't Have Nice Software](https://andrewkelley.me/post/why-we-cant-have-nice-software.html)

[LeCun on 2023-02-13](https://twitter.com/ylecun/status/1625118108082995203) (780.8K views as of 2024-02-06, [archived](https://web.archive.org/web/20230213173604/https://twitter.com/ylecun/status/1625118108082995203)):

> My unwavering opinion on current (auto-regressive) LLMs
>
> 1. They are useful as writing aids.
> 3. They make stuff up or retrieve stuff approximately.
> 6. Current LLMs should be used as writing aids, not much more.
> 7. Marrying them with tools such as search engines is highly non trivial.
> 8. There will be better systems that are factual, non toxic, and controllable. They just won't be auto-regressive LLMs.
> [...]
> 10. Warning folks that AR-LLMs make stuff up and should not be used to get factual advice.
> 11. Warning that only a small superficial portion of human knowledge can ever be captured by LLMs.
> 12. Being clear that better system will be appearing, but they will be based on different principles. They will not be auto-regressive LLMs.
> [...]

All this does not seem to stop people to build (lots of) stuff.

## LLM tooling

Proliferation of new frameworks and tool categories. Some problems:

* download, models onto machine, copy, packaging, wrapper, customization (e.g. [ollama](https://ollama.com) to run [local models](https://github.com/miku/localmodels))
* api wrappers, adding custom data to the generation process; general libraries like [llamaindex](https://www.llamaindex.ai/), with [adapters](https://llamahub.ai/), these libraries then use some paid or hosted API, like openai, claude - llama and friends; [langchain](https://www.langchain.com/), ...
* running models: [llamafile](https://github.com/Mozilla-Ocho/llamafile)

## lingoose

Many tools written in Python, but is there something similar in Go?

```
$ go run examples/hello.go
2024/02/27 14:41:42 using default ollama endpoint with model stablelm2:1.6b-zephyr-fp16
Thread:
user:
        Type: text
        Text: tell me a joke about geese
assistant:
        Type: text
        Text: Why did the chicken cross the road?

To get to the other side, of course!

But seriously, here's another one:

What do you call a group of geese with no leader?

A gaggle!
```

## Tasks

* setup threads, communicate with LLM via API (different options)
* loading content (e.g. wrappers aroung filetypes, data sources, like pubmed, ...)
* vector database (or just json file)

Two examples:

* DLF text search (dlfq)
* wikipedia outlink RAG

## DLFQ






