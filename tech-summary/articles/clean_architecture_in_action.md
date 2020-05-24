
## [A Guided Tour inside a clean architecture code base](https://proandroiddev.com/a-guided-tour-inside-a-clean-architecture-code-base-48bb5cc9fc97)

<img src="resources/clean_arch_1_3smlPZenpAtICXdgcjuHSg.jpeg" alt="clean_arch_1_3smlPZenpAtICXdgcjuHSg.jpeg" width="600"/>

### Domain layer

- `Domain entities` like `value objects`, `usecases` act as `operations`
- Abstraction of `MoviesCache`, `MoviesDataStore`, `MoviesRepository` from `domain/interface`
   + `MoviesRepository` provide functionality to support `use cases`
   + `MoviesDataStore` hides retrieve data local data access or remote access
   + `MoviesCache` abstract cache operations.  From code [here](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/presentation/src/main/kotlin/com/yossisegev/movienight/di/modules/DataModule.kt#L40-L41) you could find `cachedMoviesDataStore` is implemented based on interface from `MoviesCache` 

- `usecases`

<img src="https://miro.medium.com/max/940/1*nNYYtpoOntU-uT5xitHt_Q.png" alt="clean_arch_1_3smlPZenpAtICXdgcjuHSg.jpeg" width="600"/>

- When test `usecases` in domain layer, use `mockito` to mock result from `interface/MoviesRepository` and focus to test logic in `usecases`
```kotlin
    @Test
    fun getMovieDetailsById() {
        val movieEntity = DomainTestUtils.getTestMovieEntity(100)
        val movieRepository = Mockito.mock(MoviesRepository::class.java)
        val getMovieDetails = GetMovieDetails(TestTransformer(), movieRepository)

        Mockito.`when`(movieRepository.getMovie(100)).thenReturn(Observable.just(Optional.of(movieEntity)))

        getMovieDetails.getById(100).test()
                .assertValue { returnedMovieEntity ->
                    returnedMovieEntity.hasValue() &&
                            returnedMovieEntity.value?.id == 100
                }
                .assertComplete()
    }

```
You could find more results [here](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/domain/src/test/kotlin/com/yossisegev/domain/UseCasesTests.kt#L23)

- [`DomainTestUtils`](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/domain/src/main/kotlin/com/yossisegev/domain/common/DomainTestUtils.kt#L11) provide data for `value objects` used for testing.  

## Data layer
- Purpose
   + provide all the data the application needs to function
       * Implementation details, like [`MoviesRepositoryImpl`](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/data/src/main/kotlin/com/yossisegev/data/repositories/MoviesRepositoryImpl.kt#L13) , which can be used by domain layer's `usecases`
       * The data sources implementation is abstracted, `MoviesDataStore`
       * `Mappers` has been used for conversion between different object representation, like from HTTP's GSON format or database's ['MovieData'](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/data/src/main/kotlin/com/yossisegev/data/entities/MovieData.kt#L12) to domain format of `MovieEntity`, you could find example [here](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/data/src/main/kotlin/com/yossisegev/data/mappers/MovieDataEntityMapper.kt#L13)

- Tests
  + the implementation of MoviesRepositoryImpl, via [`MovieRepositoryImplTests`](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/data/src/test/kotlin/com/yossisegev/data/MovieRepositoryImplTests.kt#L21)
  + Functionality of mappers, [`MappersTests`](https://github.com/mrsegev/MovieNight/blob/7df52e6c93d6932b4b58de9f4f906f86df93dce1/data/src/test/kotlin/com/yossisegev/data/MappersTests.kt#L47)
  + Implementation of `MovieDataStore`, `MovieCache`, etc


## The presentation layer
Example of `PopularMoviesViewModel`
```kotlin
class PopularMoviesViewModel(private val getPopularMovies: GetPopularMovies,
                             private val movieEntityMovieMapper: Mapper<MovieEntity, Movie>):
        BaseViewModel() {

    var viewState: MutableLiveData<PopularMoviesViewState> = MutableLiveData()
    var errorState: SingleLiveEvent<Throwable?> = SingleLiveEvent()

    init {
        viewState.value = PopularMoviesViewState()
    }

    fun getPopularMovies() {
        addDisposable(getPopularMovies.observable()
                .flatMap { movieEntityMovieMapper.observable(it) }
                .subscribe({ movies ->
                    viewState.value?.let {
                        val newState = this.viewState.value?.copy(showLoading = false, movies = movies)
                        this.viewState.value = newState
                        this.errorState.value = null
                    }
                }, {
                    viewState.value = viewState.value?.copy(showLoading = false)
                    errorState.value = it
                }))
    }
}
```

## [Trying Clean Architecture on Golang](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)


<img src="resources/clean-arch_golang.png" alt="clean-arch_golang.png" width="600"/>
