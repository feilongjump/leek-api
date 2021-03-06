<?php

namespace App\Providers;

use App\Models\Article;
use App\Models\Content;
use App\Models\Project;
use App\Models\ProjectColumn;
use App\Models\ProjectColumnCard;
use App\Observers\ArticleObserver;
use App\Observers\ContentObserver;
use App\Observers\ProjectColumnCardObserver;
use App\Observers\ProjectColumnObserver;
use App\Observers\ProjectObserver;
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
        Project::observe(ProjectObserver::class);
        ProjectColumn::observe(ProjectColumnObserver::class);
        ProjectColumnCard::observe(ProjectColumnCardObserver::class);
    }
}
