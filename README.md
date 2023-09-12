# Intro to Go

<h2>Introduction</h2>
<p>
  In this repo, I have basically solved two short problems to understand Go programming language.
</p>

<p>
  <b>Q1 - Top K words:</b> The task is to find the <tt>K</tt> most common words in a
  given document. To exclude common words such as "a" and "the", the user of the program
  should be able to specify the minimum character threshold for a word. Word matching is
  case insensitive and punctuation should be removed. We can find more details on what
  qualifies as a word in the comments in the code.
</p>

<p>
  <b>Q2 - Parallel sum:</b> The task is to implement a function that sums a list of
  numbers in a file in parallel. For this problem we'll use goroutines (the
  <tt>go</tt> keyword) and channels to pass messages across the goroutines. While it is
  possible to just sum all the numbers sequentially, the point of this problem is to
  familiarize us with the synchronization mechanisms in Go.
</p>

<h3>Testing</h3>

<p>
  We'll use <tt>top_words_test.go</tt> and <tt>sum_test.go</tt> as the test functions.
  to test the correctness of the code, run the following:
</p>
<pre>
  $ go test
</pre>
<p>
  If all tests pass, you should see the following output:
</p>
<pre>
  $ go test
  PASS
  ok      
</pre>