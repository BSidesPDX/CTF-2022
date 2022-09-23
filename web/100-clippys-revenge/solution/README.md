# 100 - Clippy's Revenge (Solution)

Visiting the challenge site, a free (but wrong) flag is presented in a popup box. If you try to copy it, you will actually get a different (also wrong) flag.

If you visit the challenge site in a Chromium based browser, the flag is automatically placed on your clipboard due to a (probably) bug that allows Chromium browsers to write text to your clipboard without user interaction.

Alternatively, you may notice that the jquery library load does not have an integrity hash. If you open `/assets/jquery-3.4.1.slim.min.js` you can grep for `BSidesPDX` and find the flag.
