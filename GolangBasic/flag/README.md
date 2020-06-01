### Flag 官方有三种用法
一、Example 1: A single string flag called "species" with default value "gopher".

二、// Example 2: Two flags sharing a variable, so we can have a shorthand.
  // The order of initialization is undefined, so make sure both use the
  // same default value. They must be set up with an init function.
  
三、// Example 3: A user-defined flag type, a slice of durations.

四、flag-set, 在 nsq or jaeger repo 里面有看见这种用户，借鉴一下，用模板复制一下

 