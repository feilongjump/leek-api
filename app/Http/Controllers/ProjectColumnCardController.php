<?php

namespace App\Http\Controllers;

use App\Models\Project;
use App\Models\ProjectColumn;
use Illuminate\Http\Request;
use App\Models\ProjectColumnCard;
use App\Http\Requests\ProjectColumnCardRequest;
use App\Http\Resources\ProjectColumnCardResource;

class ProjectColumnCardController extends Controller
{
    public function __construct()
    {
        $this->middleware(['auth:sanctum']);
    }

    public function index(Request $request)
    {
        $projects = ProjectColumnCard::latest()->get();

        return ProjectColumnCardResource::collection($projects);
    }

    public function store(Project $project, ProjectColumn $column, ProjectColumnCardRequest $request)
    {
        $this->authorize('create', Project::class);

        return new ProjectColumnCardResource($column->cards()->create($request->all()));
    }

    public function show(Project $project, ProjectColumn $column, ProjectColumnCard $card)
    {
        $this->authorize('view', $card);

        $card->loadMissing('content');

        return new ProjectColumnCardResource($card);
    }

    public function update(Project $project, ProjectColumn $column, ProjectColumnCardRequest $request, ProjectColumnCard $card)
    {
        $this->authorize('update', $card);

        $card->update($request->all());

        return new ProjectColumnCardResource($card);
    }

    public function destroy(Project $project, ProjectColumn $column, ProjectColumnCard $card)
    {
        $this->authorize('delete', $card);

        $card->delete();

        return $this->withNoContent();
    }
}
