<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\SoftDeletes;

class ProjectColumn extends Model
{
    use HasFactory, SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var string[]
     */
    protected $fillable = [
        'project_id', 'name',
    ];

    public function cards()
    {
        return $this->hasMany(ProjectColumnCard::class);
    }

    public function project()
    {
        return $this->belongsTo(Project::class);
    }
}
