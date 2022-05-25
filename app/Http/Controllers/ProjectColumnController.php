<?php

namespace App\Http\Controllers;

use App\Models\Project;
use App\Models\ProjectColumn;
use App\Http\Requests\ProjectRequest;
use App\Http\Resources\ProjectColumnResource;

class ProjectColumnController extends Controller
{
    public function __construct()
    {
        $this->middleware('auth:sanctum');
    }

    public function store(Project $project, ProjectRequest $request)
    {
        $this->authorize('create', Project::class);

        return new ProjectColumnResource($project->columns()->create($request->all()));
    }

    public function update(Project $project, ProjectRequest $request, ProjectColumn $column)
    {
        $this->authorize('update', $project);

        $column->update($request->all());

        return new ProjectColumnResource($column);
    }

    public function destroy(Project $project, ProjectColumn $column)
    {
        $this->authorize('delete', $project);

        $column->delete();

        return $this->withNoContent();
    }
}
