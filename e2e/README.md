# TODO

- Should take advantage of randomized db entities. We can test their creation
  via UI, but instead of reusing a UI creation function, create data for every test via
  API calls (we can create all users with api keys so that we don't have a
  generic admin making the calls, since some endpoints will behave differently
  based on current user).

  We can generate an e2e ts client easily also with orval bare axios client:
  https://github.com/anymaniax/orval/blob/master/samples/react-app/src/api/endpoints/petstoreFromFileSpecWithTransformer.ts
