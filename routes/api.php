<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::post('login', [\App\Http\Controllers\AuthController::class, 'login']);

Route::get('me', [\App\Http\Controllers\UserController::class, 'me']);

Route::apiResources([
    'articles' => \App\Http\Controllers\ArticleController::class,
    'projects' => \App\Http\Controllers\ProjectController::class,
    'projects/{project}/columns/{column}/cards' => \App\Http\Controllers\ProjectColumnCardController::class,
]);
Route::resource(
    'projects/{project}/columns',
    \App\Http\Controllers\ProjectColumnController::class,
    ['only' => ['index', 'store', 'update', 'destroy']]
);

Route::get('novels/search', [\App\Http\Controllers\NovelController::class, 'search']);
Route::get('novels/{novel}/chapters/{chapter}', [\App\Http\Controllers\NovelController::class, 'chapter']);
Route::get('novels/{novel}', [\App\Http\Controllers\NovelController::class, 'show']);
