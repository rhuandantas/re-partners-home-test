Good README: Ensure the README file is clear and detailed.
Commented Code: Add meaningful comments to explain the code logic.
Containerized Code: Use containerization for easy deployment and testing.
Unit Tests: Include appropriate unit tests to validate functionality.
Good Code Architecture: Design the code with a solid structure. Adding a persistence layer could make the service more robust.
Working Code: Ensure the code works as expected.
Edge Case Configuration
To verify the code functionality, use the following edge case:
    Pack Sizes: 23, 31, 53
    Amount: 500,000
    Desired Output: {23: 2, 31: 7, 53: 9429}
    Additionally, a live demo of the task is highly encouraged.
Also after we are done with the task and it is approved internally I will need your resume so we can prepare that to be presented to Gymshark.
Let me know if you have any questions, I will be happy to answer!
-------------------------
He needs to ensure that his task is correct
Tell him to make sure that we are able to change the pack sizes and calculate minimizing the number of packs/quantity
Think of Money. Changing a dollar.
4 quarters equal a dollar
so does 2 dimes, 1 nickel and 3 quarters
lots of variations, but he needs to make sure that we are able to change the pack sizes and calculate
----
### 1st Review -
Readme instructions -
He should include instructions to run the container
- Code Commented
- Containerized
- Add unit test
- Code Arch is fair 
- Code is not working fine - Setting the
- Live demo to test
### 2nd Review - >> The algorithm is working
OK, however he did not include instructions to run/build the container.
Also, the API worked with the make commands, but not when I tried reaching the endpoint if it was ran in the container... 
So my final review would be to fix the container documentation and make sure the container works, besides that, I think it's OK.