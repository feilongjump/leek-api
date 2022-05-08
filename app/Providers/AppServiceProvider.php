<?php

namespace App\Providers;

use App\Models\Article;
use App\Models\Content;
use App\Observers\ArticleObserver;
use App\Observers\ContentObserver;
use Illuminate\Http\Resources\Json\JsonResource;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }

    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        JsonResource::withoutWrapping();

        Article::observe(ArticleObserver::class);
        Content::observe(ContentObserver::class);
    }
}
