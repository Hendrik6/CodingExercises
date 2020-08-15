// In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.
// If the town judge exists, then:
// The town judge trusts nobody.
// Everybody (except for the town judge) trusts the town judge.
// There is exactly one person that satisfies properties 1 and 2.
// You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.
// If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.

/**
 * @param {number} N
 * @param {number[][]} trust
 * @return {number}
 */
var findJudge = function (N, trust) {
  const truster = Array(N).fill(0);
  const trustee = Array(N).fill(0);

  for (let i = 0; i < trust.length; i++) {
    let [a, b] = trust[i];
    a--;
    b--;

    truster[a]++;
    trustee[b]++;
  }

  for (let i = 0; i < N; i++) {
    if (truster[i] == 0 && trustee[i] == N - 1) {
      return i + 1;
    }
  }

  return -1;
};
