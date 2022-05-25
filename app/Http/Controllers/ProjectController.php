<?php

namespace App\Http\Controllers;

use App\Models\Project;
use Illuminate\Http\Request;
use App\Http\Requests\ProjectRequest;
use App\Http\Resources\ProjectResource;

class ProjectController extends Controller
{
    public function __construct()
    {
        $this->middleware(['auth:sanctum']);
    }

    public function index(Request $request)
    {
        $projects = Project::whereUserId(auth()->user()->id)
            ->latest()
            ->paginate($request->get('per_page', 10));

        return ProjectResource::collection($projects);
    }

    public function store(ProjectRequest $request)
    {
        $this->authorize('create', Project::class);

        return new ProjectResource(Project::create($request->all()));
    }

    public function show(Project $project)
    {
        $this->authorize('view', $project);

        $project->loadMissing(['content', 'columns']);

        return new ProjectResource($project);
    }

    public function update(ProjectRequest $request, Project $project)
    {
        $this->authorize('update', $project);

        $project->update($request->all());

        return new ProjectResource($project);
    }

    public function destroy(Project $project)
    {
        $this->authorize('delete', $project);

        $project->delete();

        return $this->withNoContent();
    }
}
