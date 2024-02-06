# An example RAG project using a vector database

> [Ouroboros](https://en.wikipedia.org/wiki/Ouroboros)

![](static/Chrysopoea.png)

An LLM is unreliable.

> Have a peek at this blog post that is going around lately: [The pain points
> of building a copilot](https://austinhenley.com/blog/copilotpainpoints.html)
> These people are brimming with excitement about all the new problems that
> LLMs are bringing to the table. -- [Why We Can't Have Nice Software](https://andrewkelley.me/post/why-we-cant-have-nice-software.html)

[LeCun on 2023-02-13](https://twitter.com/ylecun/status/1625118108082995203) (780.8K views as of 2024-02-06):

> My unwavering opinion on current (auto-regressive) LLMs
> 1. They are useful as writing aids.
> 2. They are "reactive" & don't plan nor reason.
> 3. They make stuff up or retrieve stuff approximately.
> 4. That can be mitigated but not fixed by human feedback.
> 5. Better systems will come
> 6. Current LLMs should be used as writing aids, not much more.
> 7. Marrying them with tools such as search engines is highly non trivial.
> 8. There *will* be better systems that are factual, non toxic, and controllable. They just won't be auto-regressive LLMs.
> [...]
> 10. Warning folks that AR-LLMs make stuff up and should not be used to get factual advice.
> 11. Warning that only a small superficial portion of human knowledge can ever be captured by LLMs.
> 12. Being clear that better system will be appearing, but they will be based on different principles. They will not be auto-regressive LLMs.
> 13. Why do LLMs appear much better at generating code than generating general text? Because, unlike the real world, the universe that a program manipulates (the state of the variables) is limited, discrete, deterministic, and fully observable. The real world is none of that
> 14. Unlike what the most acerbic critics of Galactica have claimed: LLMs *are* being used as writing aids.; They *will not* destroy the fabric of society by causing the mindless masses to believe their made-up nonsense.; People will use them for what they are helpful with.
